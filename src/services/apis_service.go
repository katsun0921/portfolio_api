package services

import (
	"github.com/katsun0921/go_utils/rest_errors"
	"github.com/katsun0921/portfolio_api/src/constants"
	"github.com/katsun0921/portfolio_api/src/domain/apis"
	"regexp"
	"sort"
	"strings"
	"time"
)

var (
	ApisService apisServiceInterface = &apisService{}
)

type apisServiceInterface interface {
	GetApiAll() ([]*apis.Api, rest_errors.RestErr)
	GetRss(service string) ([]*apis.Api, rest_errors.RestErr)
	GetTwitter() ([]*apis.Api, rest_errors.RestErr)
	GetSkills() ([]apis.Skill, rest_errors.RestErr)
	GetWorkExpress() ([]apis.WorkExpress, rest_errors.RestErr)
}

type apisService struct {
}

func (api *apisService) GetApiAll() ([]*apis.Api, rest_errors.RestErr) {
	var res []*apis.Api

	tweets, errTwitter := api.GetTwitter()
	if errTwitter != nil {
		return nil, errTwitter
	}
	zenn, errZenn := api.GetRss(constants.ZENN)
	if errZenn != nil {
		return nil, errZenn
	}
	res = append(res, tweets...)
	res = append(res, zenn...)
	sort.SliceStable(res, func(i, j int) bool { return res[i].DateUnix > res[j].DateUnix })
	return res, nil
}

func (*apisService) GetRss(service string) ([]*apis.Api, rest_errors.RestErr) {
	api := &apis.Api{}
	//article := &articles.Article{}
	var res []*apis.Api
	feed, err := api.GetFeedApi(service)
	if err != nil {
		return nil, err
	}

	/* TODO: Comment out until post is made.
	articleId, articleErr := article.FindByLatestArticleId(service)
	if articleErr != nil {
	  return nil, articleErr
	}
	*/

	items := feed.Items

	for _, item := range items {

		/* TODO: Comment out until post is made.
		   if item.GUID == articleId {
		     break
		   }
		*/
		key := &apis.Api{}
		itemPlainText := item.Description
		itemPlainText = strings.ReplaceAll(itemPlainText, " ", "")
		itemPlainText = strings.ReplaceAll(itemPlainText, "\n", "")
		t, _ := time.Parse(time.RFC1123, feed.Published)
		feedDate := t.Format(constants.DateLayout)

		key.Id = item.GUID
		key.Text = item.Title + "\n" + itemPlainText
		key.Link = item.Link
		key.DateCreated = feedDate
		key.DateUnix = int(t.Unix())
		key.Service = service

		res = append(res, key)
	}

	return res, nil
}

func (*apisService) GetTwitter() ([]*apis.Api, rest_errors.RestErr) {
	api := &apis.Api{}
	var res []*apis.Api
	tweets, err := api.GetTwitterApi()
	if err != nil {
		return nil, err
	}

	for _, tweet := range tweets {
		isRetweeted := tweet.Retweeted
		if isRetweeted {
			continue
		}

		key := &apis.Api{}
		tweetText := strings.ReplaceAll(tweet.Text, "\n", "")
		regLink := regexp.MustCompile("https://.*$")
		tweetPlainText := regLink.ReplaceAllString(tweetText, "")
		tweetPlainText = strings.TrimSpace(tweetPlainText)
		tweetScreenName := tweet.User.ScreenName
		tweetStatus := tweet.IDStr
		tweetLink := constants.TwitterDomain + "/" + tweetScreenName + "/status/" + tweetStatus

		t, _ := time.Parse(time.UnixDate, tweet.CreatedAt)
		tweetDate := t.Format(constants.DateLayout)
		key.Id = tweet.IDStr
		key.Text = tweetPlainText
		key.Link = tweetLink
		key.DateCreated = tweetDate
		key.DateUnix = int(t.Unix())
		key.Service = constants.TWITTER

		res = append(res, key)
	}

	return res, nil
}

func (*apisService) GetSkills() ([]apis.Skill, rest_errors.RestErr) {
	api := &apis.Api{}
	var res []apis.Skill
	skills, skillsErr := api.GetGoogleSheetsApi(constants.SheetRangeSkill)
	if skillsErr != nil {
		return nil, skillsErr
	}

	jobNames, jobNamesErr := api.GetGoogleSheetsApi(constants.SheetRangeJobType)
	if jobNamesErr != nil {
		return nil, skillsErr
	}

	for _, skill := range skills {
		job := apis.Skill{}
		lang := apis.Language{}
		if name, ok := skill[2].(string); ok {
			lang.Name = name
		}

		if level, ok := skill[3].(string); ok {
			lang.Level = level
		}

		if jobType, ok := skill[1].(string); ok {
			job.Job = jobType
		}

		var arrLang []apis.Language
		var arrJobs []apis.Skill
		arrLang = append(arrLang, lang)

		isJobType := false
		for i := 0; i < len(res); i++ {
			if res[i].Job == job.Job {
				res[i].Skills = append(res[i].Skills, arrLang...)
				isJobType = true
				break
			}
		}

		if !isJobType {
			switch job.Job {
			case constants.Frontend:
				job.Id = "1"
				job.Job = getJobName(constants.Frontend, jobNames)
			case constants.Backend:
				job.Id = "2"
				job.Job = getJobName(constants.Backend, jobNames)
			case constants.Infra:
				job.Id = "3"
				job.Job = getJobName(constants.Infra, jobNames)
			}

			job.Skills = arrLang
			arrJobs = append(arrJobs, job)
			res = append(res, arrJobs...)
			continue
		}
	}

	return res, nil
}

func getJobName(jobType string, jobNames [][]interface{}) string {
	for _, jobName := range jobNames {
		if jobType == jobName[0] {
			if jobName, ok := jobName[1].(string); ok {
				return jobName
			}
		}
	}
	return jobType
}

func (*apisService) GetWorkExpress() ([]apis.WorkExpress, rest_errors.RestErr) {
	api := &apis.Api{}
	var res []apis.WorkExpress
	workExpress, workExpressErr := api.GetGoogleSheetsApi(constants.SheetRangeWorkExpress)
	if workExpressErr != nil {
		return nil, workExpressErr
	}

	for _, work := range workExpress {
		express := apis.WorkExpress{}
		if company, ok := work[0].(string); ok {
			express.Company = company
		}

		if project, ok := work[1].(string); ok {
			express.Project = project
		}

		if jobType, ok := work[2].(string); ok {
			express.JobType = jobType
		}

		if startDate, ok := work[3].(string); ok {
			express.StartDate = startDate
		}

		if endDate, ok := work[4].(string); ok {
			express.EndDate = endDate
		}

		if description, ok := work[5].(string); ok {
			express.Description = description
		}

		res = append(res, express)
	}

	for _, _res := range res {
		setUnixParse(_res.EndDate)
	}

	sort.SliceStable(res, func(i, j int) bool { return setUnixParse(res[i].EndDate)  > setUnixParse(res[j].EndDate) })

	return res, nil
}

func setUnixParse(date string) int {
	t, _ := time.Parse(constants.SheetDateLayout, date)
	return int(t.Unix())
}
