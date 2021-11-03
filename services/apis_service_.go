package services

import (
	"github.com/katsun0921/go_utils/rest_errors"
	"github.com/katsun0921/portfolio_api/domain/apis"
	"strings"
)

const (
	RSS  string = "RSS"
	ZENN string = "zenn"
)

var (
	ApisService apisServiceInterface = &apisService{}
)

type apisServiceInterface interface {
	GetApi(api apis.Api) (*apis.Api, rest_errors.RestErr)
}

type apisService struct {
}

func (a *apisService) GetApi(api apis.Api) (*apis.Api, rest_errors.RestErr) {
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
		result.Description.PlainText = itemPlainText
		result.Description.Html = item.Description
		result.Link = item.Link
		result.DateCreated = item.Published
		result.Type = RSS
		result.Service = ZENN
	}

	return result, nil
}
