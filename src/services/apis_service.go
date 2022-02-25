package services

import (
	"github.com/katsun0921/go_utils/rest_errors"
	"github.com/katsun0921/portfolio_api/src/constants"
	"github.com/katsun0921/portfolio_api/src/domain/apis"
	"github.com/katsun0921/portfolio_api/src/utils/date_utils"
	"regexp"
	"sort"
	"strconv"
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
	GetWorkexpress() ([]apis.Workexpress, rest_errors.RestErr)
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
		key.DateUnix = date_utils.SetUnixParse(time.UnixDate, tweet.CreatedAt)
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

	jobTypes, jobTypesErr := api.GetGoogleSheetsApi(constants.SheetRangeJobType)
	if jobTypesErr != nil {
		return nil, skillsErr
	}

	for _, skill := range skills {
		job := apis.Skill{}
		lang := apis.Programming{}

		if jobId, ok := skill[0].(string); ok {
			lang.Id = jobId
		}

		if language, ok := skill[2].(string); ok {
			lang.Language = language
		}

		if level, ok := skill[3].(string); ok {
			lang.Level = level
		}

		if jobType, ok := skill[1].(string); ok {
			job.Name = jobType
			job.NameJp = jobType
		}

		var arrLang []apis.Programming
		var arrJobs []apis.Skill
		arrLang = append(arrLang, lang)

		isJobType := false
		for i := 0; i < len(res); i++ {
			if res[i].Name == job.Name {
				res[i].Programming = append(res[i].Programming, arrLang...)
				isJobType = true
				break
			}
		}

		if !isJobType {
			job.Programming = arrLang
			arrJobs = append(arrJobs, job)
			res = append(res, arrJobs...)
			continue
		}
	}

	// Set id and NameJp
	for i := 0; i < len(res); i++ {
		switch res[i].NameJp {
		case constants.Frontend:
			res[i].Id = getSkillId(constants.Frontend, jobTypes)
			res[i].NameJp = getSkillNameJp(constants.Frontend, jobTypes)
		case constants.Backend:
			res[i].Id = getSkillId(constants.Backend, jobTypes)
			res[i].NameJp = getSkillNameJp(constants.Backend, jobTypes)
		case constants.Infra:
			res[i].Id = getSkillId(constants.Infra, jobTypes)
			res[i].NameJp = getSkillNameJp(constants.Infra, jobTypes)
		}
	}

	return res, nil
}

func getSkillNameJp(resType string, jobTypes [][]interface{}) string {
	for _, jobType := range jobTypes {
		if resType == jobType[1] {
			if name, ok := jobType[2].(string); ok {
				return name
			}
		}
	}
	return resType
}

func getSkillId(resType string, jobTypes [][]interface{}) string {
	for _, jobType := range jobTypes {
		if resType == jobType[1] {
			if jobId, ok := jobType[0].(string); ok {
				return jobId
			}
		}
	}
	return ""
}

func (*apisService) GetWorkexpress() ([]apis.Workexpress, rest_errors.RestErr) {
	api := &apis.Api{}
	var res []apis.Workexpress
	Workexpress, WorkexpressErr := api.GetGoogleSheetsApi(constants.SheetRangeWorkexpress)
	if WorkexpressErr != nil {
		return nil, WorkexpressErr
	}

	for _, work := range Workexpress {
		express := apis.Workexpress{}
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

		if skills, ok := work[6].(string); ok {
			express.Skills = strings.Split(skills, ",")
		}
		res = append(res, express)
	}

	sort.SliceStable(res, func(i, j int) bool { return date_utils.SetUnixParse(constants.SheetDateLayout, res[i].EndDate) > date_utils.SetUnixParse(constants.SheetDateLayout, res[j].EndDate) })

	for i := 0; i < len(res); i++ {
		res[i].Id = strconv.Itoa(i + 1)
	}

	return res, nil
}
