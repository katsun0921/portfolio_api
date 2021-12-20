package articles

import (
	"errors"
	"fmt"
  "github.com/katsun0921/portfolio_api/src/datasources/mysql/articles_db"
	"github.com/katsun0921/go_utils/rest_errors"
	"github.com/katsun0921/go_utils/logger"
)

const (
	queryInsertArticle             = "INSERT INTO articles(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetArticle                = "SELECT id, first_name, last_name, email, date_created, status FROM articles WHERE id=?;"
	queryFindByType           = "SELECT id, first_name, last_name, email, date_created, status FROM articles WHERE status=?;"
)

func (article *Article) Get() rest_errors.RestErr {
	stmt, err := articles_db.Client.Prepare(queryGetArticle)
	if err != nil {
		logger.Error("error when trying to prepare get article statement", err)
		return rest_errors.NewInternalServerError("error when trying to get article", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(article.Id)

	if getErr := result.Scan(&article.Id, &article.FirstName, &article.LastName, &article.Email, &article.DateCreated, &article.Type); getErr != nil {
		logger.Error("error when trying to get article by id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get article", errors.New("database error"))
	}

	return nil
}

func (article *Article) Save() rest_errors.RestErr {
	stmt, err := articles_db.Client.Prepare(queryInsertArticle)
	if err != nil {
		logger.Error("error when trying to prepare save article statement", err)
		return rest_errors.NewInternalServerError("error when trying to save article", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(article.FirstName, article.LastName, article.Email, article.DateCreated, article.Type, article.Password)
	if saveErr != nil {
		logger.Error("error when trying to save article", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save article", errors.New("database error"))
	}

	articleId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert creating a new article", err)
		return rest_errors.NewInternalServerError("error when trying to save article", errors.New("database error"))
	}

	article.Id = articleId
	return nil
}


func (article *Article) FindByType(status string) ([]Article, rest_errors.RestErr) {
	stmt, err := articles_db.Client.Prepare(queryFindByType)
	if err != nil {
		logger.Error("error when trying to prepare find articles by status statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to find by status", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find articles by status", err)
		return nil, rest_errors.NewInternalServerError("error when trying to find by status", errors.New("database error"))
	}
	defer rows.Close()

	results := make([]Article, 0)
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.Id, &article.FirstName, &article.LastName, &article.Email, &article.DateCreated, &article.Type); err != nil {
			logger.Error("error when scan article row into article struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to find by status", errors.New("database error"))
		}
		results = append(results, article)
	}

	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no articles matching status %s", status), errors.New("database error"))
	}
	return results, nil
}
