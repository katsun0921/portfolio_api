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
  var res []*apis.Api
  rss, err := api.GetFeedApi(service)
  if err != nil {
    return nil, err
  }

  feeds := rss.Items
  for i := 0; i < constants.ArticlesMaxCount; i++ {
    if i >= len(feeds) {
      break
    }
    key := &apis.Api{}
    itemPlainText := feeds[i].Description
    itemPlainText = strings.ReplaceAll(itemPlainText, " ", "")
    itemPlainText = strings.ReplaceAll(itemPlainText, "\n", "")
    t, _ := time.Parse(constants.TimeLayoutRFC1123, feeds[i].Published)
    feedDate := t.Format(constants.DateLayout)

    key.Id = feeds[i].GUID
    key.Text = feeds[i].Title + "\n" + itemPlainText
    key.Link = feeds[i].Link
    key.DateCreated = feedDate
    key.DateUnix = int(t.Unix())
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

  i := 0
  maxCount := constants.ArticlesMaxCount
  for i < maxCount {
    if i >= len(tweets) {
      break
    }
    isRetweeted := tweets[i].Retweeted; if isRetweeted {
      i++
      maxCount++
      continue
    }

    key := &apis.Api{}
    tweetText := strings.ReplaceAll(tweets[i].Text, "\n", "")
    regLink := regexp.MustCompile("https://.*$")
    tweetPlainText := regLink.ReplaceAllString(tweetText, "")
    tweetPlainText = strings.TrimSpace(tweetPlainText)
    tweetScreenName := tweets[i].User.ScreenName
    tweetStatus := tweets[i].IDStr
    tweetLink := constants.TwitterDomain + "/" + tweetScreenName + "/status/" + tweetStatus

    t, _ := time.Parse(constants.TimeLayoutUnixDate, tweets[i].CreatedAt)
    tweetDate := t.Format(constants.DateLayout)
    key.Id = tweets[i].IDStr
    key.Text = tweetPlainText
    key.Link = tweetLink
    key.DateCreated = tweetDate
    key.DateUnix = int(t.Unix())
    key.Service = constants.TWITTER

    res = append(res, key)
    i++
  }

  return res, nil
}
