package articles

import (
  "github.com/katsun0921/go_utils/rest_errors"
  "strings"
)

type Article struct {
	Id          int64  `json:"id"`
	Text        string `json:"text"`
	Link        string `json:"link"`
	Service     string `json:"service"`
	ArticleId   string `json:"article_id"`
	DateCreated string `json:"date_created"`
}

type Articles []Article

func (article *Article) Validate() rest_errors.RestErr {
	article.Text = strings.TrimSpace(article.Text)
	article.Link = strings.TrimSpace(article.Link)

	return nil
}
