package services

import (
  "fmt"
  "github.com/katsun0921/go_utils/rest_errors"
  "github.com/katsun0921/portfolio_api/domain/apis"
  "strings"
)

const (
  ZENN string = "zenn"
  TWITTER string = "twitter"
)

var (
  ApisService apisServiceInterface = &apisService{}
)

type apisServiceInterface interface {
  GetApi() (*apis.Api, rest_errors.RestErr)
  GetTwitterApi() (*apis.Api, rest_errors.RestErr)
}

type apisService struct {
}

func (*apisService) GetApi() (*apis.Api, rest_errors.RestErr) {
  result := &apis.Api{}
  feed, err := result.GetRss()
  if err != nil {
    return nil, err
  }

  items := feed.Items
  for _, item := range items {
    itemPlainText := item.Description
    itemPlainText = strings.ReplaceAll(itemPlainText, " ", "")
    itemPlainText = strings.ReplaceAll(itemPlainText, "\n", "")
    result.Title = item.Title
    result.PlainText = itemPlainText
    result.Link = item.Link
    result.DateCreated = item.Published
    result.Service = ZENN
  }

  return result, err
}

func (*apisService) GetTwitterApi() (*apis.Api, rest_errors.RestErr) {
  result := &apis.Api{}
  feed, err := result.GetRss()
  if err != nil {
    return nil, err
  }

  twitter, err := result.GetTwitter()
  fmt.Println(twitter)
  items := feed.Items
  for _, item := range items {
    itemPlainText := item.Description
    itemPlainText = strings.ReplaceAll(itemPlainText, " ", "")
    itemPlainText = strings.ReplaceAll(itemPlainText, "\n", "")
    result.Title = item.Title
    result.PlainText = itemPlainText
    result.Link = item.Link
    result.DateCreated = item.Published
    result.Service = TWITTER
  }

  return result, err
}
