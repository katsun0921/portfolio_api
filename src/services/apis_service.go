package services

import (
  "github.com/katsun0921/go_utils/rest_errors"
  "github.com/katsun0921/portfolio_api/src/constants"
  "github.com/katsun0921/portfolio_api/src/domain/apis"
  "github.com/katsun0921/portfolio_api/src/domain/articles"
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
}

type apisService struct {
}

func (api *apisService) GetApiAll() ([]*apis.Api, rest_errors.RestErr) {
  var res []*apis.Api

  tweets, errTwitter := api.GetTwitter()
  if errTwitter != nil {
    return nil, errTwitter
  }
  zenns, errZenns := api.GetRss(constants.ZENN)
  if errZenns != nil {
    return nil, errZenns
  }
  res = append(res, tweets...)
  res = append(res, zenns...)
  sort.SliceStable(res, func(i, j int) bool { return res[i].DateUnix > res[j].DateUnix })
  return res, nil
}

func (*apisService) GetRss(service string) ([]*apis.Api, rest_errors.RestErr) {
  api := &apis.Api{}
  article := &articles.Article{}
  var res []*apis.Api
  rss, err := api.GetFeedApi(service)
  if err != nil {
    return nil, err
  }

  articleId, articleErr := article.FindByLatestArticleId(service)
  if articleErr != nil {
    return nil, articleErr
  }

  feeds := rss.Items
  for _, feed := range feeds {

    if feed.GUID == articleId {
      break
    }
    key := &apis.Api{}
    itemPlainText := feed.Description
    itemPlainText = strings.ReplaceAll(itemPlainText, " ", "")
    itemPlainText = strings.ReplaceAll(itemPlainText, "\n", "")
    t, _ := time.Parse(constants.TimeLayoutRFC1123, feed.Published)
    feedDate := t.Format(constants.DateLayout)

    key.Id = feed.GUID
    key.Text = feed.Title + "\n" + itemPlainText
    key.Link = feed.Link
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
