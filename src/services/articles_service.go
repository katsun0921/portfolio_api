package services

import (
  "github.com/katsun0921/go_utils/rest_errors"
  "github.com/katsun0921/portfolio_api/src/domain/articles"
  "github.com/katsun0921/portfolio_api/src/utils/date_utils"
)

var (
  ArticlesService articlesServiceInterface = &articlesService{}
)

type articlesService struct {
}

type (
  articlesServiceInterface interface {
    GetArticle(int64) (*articles.Article, rest_errors.RestErr)
    CreateArticle(articles.Article) (*articles.Article, rest_errors.RestErr)
    SearchArticle(string) ([]articles.Article, rest_errors.RestErr)
  }
)

func (s *articlesService) GetArticle(articleId int64) (*articles.Article, rest_errors.RestErr) {
  result := &articles.Article{Id: articleId}
  if err := result.Get(); err != nil {
    return nil, err
  }
  return result, nil
}

func (s *articlesService) CreateArticle(Article articles.Article) (*articles.Article, rest_errors.RestErr) {
  if err := Article.Validate(); err != nil {
    return nil, err
  }

  Article.DateCreated = date_utils.GetNowDBFormat()
  if err := Article.Save(); err != nil {
    return nil, err
  }

  return &Article, nil
}

func (s *articlesService) SearchArticle(status string) ([]articles.Article, rest_errors.RestErr) {
  dao := &articles.Article{}
  return dao.FindByStatus(status)
}
