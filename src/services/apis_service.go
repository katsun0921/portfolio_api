package services

import (
	"fmt"
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
	GetSkills() ([]*apis.Skill, rest_errors.RestErr)
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
		t, _ := time.Parse(constants.TimeLayoutRFC1123, feed.Published)
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

		t, _ := time.Parse(constants.TimeLayoutUnixDate, tweet.CreatedAt)
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

func (*apisService) GetSkills() ([]*apis.Skill, rest_errors.RestErr) {
	skill := &apis.Skill{}
	var res []*apis.Skill
	skills, err := skill.GetGoogleSheetsApi()
	if err != nil {
		return nil, err
	}

	fmt.Println(skills)

	return res, nil
}
