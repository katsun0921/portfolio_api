package services

import (
	"github.com/katsun0921/go_utils/rest_errors"
	"github.com/katsun0921/portfolio_api/src/domain/apis"
	"github.com/katsun0921/portfolio_api/src/domain/articles"
	"github.com/katsun0921/portfolio_api/src/utils/date_utils"
)

var (
	ArticlesService articlesServiceInterface = &articlesService{}
)

type TResArticle struct {
  Id          int    `json:"id"`
  Text        string `json:"text"`
  Link        string `json:"link"`
  Service     string `json:"service"`
  DateCreated string `json:"date_created"`
}


type articlesService struct {
}

type (
	articlesServiceInterface interface {
		GetArticleAll() ([]TResArticle, rest_errors.RestErr)
		CreateArticle(articles.Article, *apis.Api) (*articles.Article, rest_errors.RestErr)
	}
)

func (s *articlesService) GetArticleAll() ([]TResArticle, rest_errors.RestErr) {
	result := &articles.Article{}
	var res TResArticle
	resArticles, err := result.Get()
	if err != nil {
		return nil, err
	}
	articleAll := make([]TResArticle, 0)
	for i, article := range resArticles {
		res.Id = i
		res.Text = article.Text
		res.Link = article.Link
    res.Service = article.Service
		res.DateCreated = article.DateCreated
		articleAll = append(articleAll, res)
	}
	return articleAll, nil
}

func (s *articlesService) CreateArticle(Article articles.Article, api *apis.Api) (*articles.Article, rest_errors.RestErr) {

	if err := Article.Validate(); err != nil {
		return nil, err
	}

	Article.Text = api.Text
	Article.Link = api.Link
	Article.Service = api.Service
	Article.ArticleId = api.Id
	Article.DateCreated = api.DateCreated
	Article.CreatedAt = date_utils.GetNowDBFormat()
	if err := Article.Save(); err != nil {
		return nil, err
	}

	return &Article, nil
}
