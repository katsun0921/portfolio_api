package services

import (
  "fmt"
  "github.com/katsun0921/go_utils/rest_errors"
  "github.com/katsun0921/portfolio_api/src/constants"
  "github.com/katsun0921/portfolio_api/src/domain/apis"
  "regexp"
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

func (*apisService) GetApiAll() ([]*apis.Api, rest_errors.RestErr) {
  api := &apis.Api{}
  var res []*apis.Api

  res = append(res, api)
  return res, nil
}

func (*apisService) GetRss(service string) ([]*apis.Api, rest_errors.RestErr) {
  api := &apis.Api{}
  var res []*apis.Api
  rss, err := api.GetFeedApi(service)
  if err != nil {
    return nil, err
  }

  feeds := rss.Items
  for i, feed := range feeds {
    key := &apis.Api{}
    itemPlainText := feed.Description
    itemPlainText = strings.ReplaceAll(itemPlainText, " ", "")
    itemPlainText = strings.ReplaceAll(itemPlainText, "\n", "")
    t, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", feed.Published)
    fmt.Println(feed.Published)
    feedDate := t.Format(constants.DateLayout)

    key.Id = i
    key.Text = feed.Title + "\n" + itemPlainText
    key.Link = feed.Link
    key.DateCreated = feedDate
    key.Service = constants.ZENN

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

  for i, tweet := range tweets {
    key := &apis.Api{}
    tweetText := strings.ReplaceAll(tweet.Text, "\n", "")
    regLink := regexp.MustCompile("https:.*$")
    tweetPlainText := regLink.ReplaceAllString(tweetText,"")
    tweetPlainText = strings.TrimSpace(tweetPlainText)
    tweetLink := regLink.FindString(tweetText)

    t, _ := time.Parse("Mon Jan 2 15:04:05 MST 2006", tweet.CreatedAt)
    tweetDate := t.Format(constants.DateLayout)
    key.Id = i
    key.Text = tweetPlainText
    key.Link = tweetLink
    key.DateCreated = tweetDate
    key.Service = constants.TWITTER

    res = append(res, key)
  }

  return res, nil
}
