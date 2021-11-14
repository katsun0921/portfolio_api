package services

import (
  "fmt"
  "github.com/katsun0921/go_utils/rest_errors"
  "github.com/katsun0921/portfolio_api/src/domain/apis"
  "regexp"
  "strings"
  "time"
)

const (
  ZENN string = "zenn"
  TWITTER string = "twitter"
  layout = "2006/01/02 15:00"
)

var (
  ApisService apisServiceInterface = &apisService{}
)

type apisServiceInterface interface {
  GetApi() ([]*apis.Api, rest_errors.RestErr)
  GetTwitterApi() ([]*apis.Api, rest_errors.RestErr)
}

type apisService struct {
}

func (*apisService) GetApi() ([]*apis.Api, rest_errors.RestErr) {
  api := &apis.Api{}
  var res []*apis.Api
  rss, err := api.GetRss()
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
    feedDate := t.Format(layout)

    key.Id = i
    key.Text = feed.Title + "\n" + itemPlainText
    key.Link = feed.Link
    key.DateCreated = feedDate
    key.Service = ZENN

    res = append(res, key)
  }

  return res, nil
}

func (*apisService) GetTwitterApi() ([]*apis.Api, rest_errors.RestErr) {
  api := &apis.Api{}
  var res []*apis.Api
  tweets, err := api.GetTwitter()
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
    tweetDate := t.Format(layout)
    key.Id = i
    key.Text = tweetPlainText
    key.Link = tweetLink
    key.DateCreated = tweetDate
    key.Service = TWITTER

    res = append(res, key)
  }

  return res, nil
}
