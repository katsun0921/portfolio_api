package articles

import (
  "errors"
  "github.com/katsun0921/go_utils/rest_errors"
  "strings"
)

const (
  TypeTwitter = "twitter"
  TypeZenn    = "zenn"
)

type Article struct {
  Id          int64  `json:"id"`
  FirstName   string `json:"first_name"`
  LastName    string `json:"last_name"`
  Email       string `json:"email"`
  DateCreated string `json:"date_created"`
  Status      string `json:"status"`
  Type        string `json:"type"`
  Password    string `json:"password"`
}

type Articles []Article

func (article *Article) Validate() rest_errors.RestErr {
  article.FirstName = strings.TrimSpace(article.FirstName)
  article.LastName = strings.TrimSpace(article.LastName)

  article.Email = strings.TrimSpace(strings.ToLower(article.Email))
  if article.Email == "" {
    return rest_errors.NewBadRequestError("invalid email address", errors.New("database error"))
  }

  article.Password = strings.TrimSpace(article.Password)
  if article.Password == "" {
    return rest_errors.NewBadRequestError("invalid password", errors.New("database error"))
  }
  return nil
}
