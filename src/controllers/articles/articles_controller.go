package articles

import (
  "errors"
  "github.com/gin-gonic/gin"
  "github.com/katsun0921/go_utils/rest_errors"
  "github.com/katsun0921/portfolio_api/src/domain/apis"
  "github.com/katsun0921/portfolio_api/src/domain/articles"
  "github.com/katsun0921/portfolio_api/src/services"
  "net/http"
)

func Create(c *gin.Context) {
  var article articles.Article
  var resApis []*apis.Api

  resApis, _ = services.ApisService.GetApiAll()
  results := make([]*articles.Article, 0)

  if err := c.ShouldBindJSON(&article); err != nil {
    restErr := rest_errors.NewBadRequestError("invalid json error for articles", errors.New("json error"))
    c.JSON(restErr.Status(), restErr)
    return
  }

  for _, api := range resApis {
    saveArticle, saveErr := services.ArticlesService.CreateArticle(article, api)
    if saveErr != nil {
      c.JSON(saveErr.Status(), saveErr)
      //TODO: Handle user creation error
      return
    }
    results = append(results, saveArticle)
  }

  c.JSON(http.StatusCreated, results)
}
