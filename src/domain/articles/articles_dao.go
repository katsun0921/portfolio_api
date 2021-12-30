package articles

import (
	"errors"
	"fmt"
	"github.com/katsun0921/go_utils/logger"
	"github.com/katsun0921/go_utils/rest_errors"
	"github.com/katsun0921/portfolio_api/src/datasources/mysql/blog_db"
)

const (
	queryInsertArticle         = "INSERT INTO articles(text, link, service, article_id, data_created, created_at) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetArticle            = "SELECT id, text, link, service, article_id, created_at FROM articles WHERE id=?;"
	queryFindByService         = "SELECT * FROM articles WHERE service=?;"
	queryFindByLatestArticleId = "SELECT id, service, article_id, data_created FROM blog_db.articles WHERE data_created = (SELECT MAX(data_created) FROM blog_db.articles AS us WHERE blog_db.articles.service=?);"
)

func (article *Article) Get() rest_errors.RestErr {
	stmt, err := blog_db.Client.Prepare(queryGetArticle)
	if err != nil {
		logger.Error("error when trying to prepare get article statement", err)
		return rest_errors.NewInternalServerError("error when trying to get article", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(article.Id)

	if getErr := result.Scan(&article.Id, &article.Text, &article.Link, &article.Service, &article.ArticleId, &article.DateCreated); getErr != nil {
		logger.Error("error when trying to get article by id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get article", errors.New("database error"))
	}

	return nil
}

func (article *Article) Save() rest_errors.RestErr {
	stmt, err := blog_db.Client.Prepare(queryInsertArticle)
	if err != nil {
		logger.Error("error when trying to prepare save article statement", err)
		return rest_errors.NewInternalServerError("error when trying to connect mysql", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(article.Text, article.Link, article.Service, article.ArticleId, article.DateCreated, article.CreatedAt)
	if saveErr != nil {
		logger.Error("error when trying to save article", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save article", errors.New("database error"))
	}

	articleId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert creating a new article", err)
		return rest_errors.NewInternalServerError("error when trying to save article_id", errors.New("database error"))
	}

	article.Id = articleId
	return nil
}

func (article *Article) FindByService(service string) ([]Article, rest_errors.RestErr) {
	stmt, err := blog_db.Client.Prepare(queryFindByService)
	if err != nil {
		logger.Error("error when trying to prepare find blog_db by status statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to find by status", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(service)
	if err != nil {
		logger.Error("error when trying to find blog_db by status", err)
		return nil, rest_errors.NewInternalServerError("error when trying to find by status", errors.New("database error"))
	}
	defer rows.Close()

	results := make([]Article, 0)
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.Id, &article.Text, &article.Link, &article.Service, &article.ArticleId, &article.DateCreated, &article.CreatedAt); err != nil {
			logger.Error("error when scan article row into blog_db struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to find by status", errors.New("database error"))
		}
		results = append(results, article)
	}

	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", err), errors.New("database error"))
	}
	return results, nil
}

func (article *Article) FindByLatestArticleId(service string) (string, rest_errors.RestErr) {
  articleId := ""

	stmt, err := blog_db.Client.Prepare(queryFindByLatestArticleId)
	if err != nil {
		logger.Error("error when trying to prepare find blog_db by status statement", err)
		return articleId, rest_errors.NewInternalServerError("error when trying to find by status", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(service)
	if err != nil {
		logger.Error("error when trying to find blog_db by status", err)
		return articleId, rest_errors.NewInternalServerError("error when trying to find by status", errors.New("database error"))
	}
	defer rows.Close()

	results := make([]Article, 0)
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.Id, &article.Service, &article.ArticleId, &article.DateCreated); err != nil {
			logger.Error("error when scan article row into blog_db struct", err)
			return articleId, rest_errors.NewInternalServerError("error when trying to find by status", errors.New("database error"))
		}
		results = append(results, article)
	}

	if len(results) > 0 {
	  articleId = results[0].ArticleId
  }

	return articleId, nil
}
