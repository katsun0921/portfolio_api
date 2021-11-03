package apis

import (
  "github.com/gin-gonic/gin"
  "github.com/katsun0921/portfolio_api/domain/apis"
  "github.com/katsun0921/portfolio_api/services"
  "net/http"
)

func Get(c *gin.Context) {
  var api apis.Api

  rss, err := services.ApisService.GetApi(api)
  if err != nil {
    return
  }
  c.JSON(http.StatusOK, rss)
}
