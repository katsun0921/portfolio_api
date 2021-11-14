package apis

import (
  "errors"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/katsun0921/go_utils/logger"
  "github.com/katsun0921/go_utils/rest_errors"
  "github.com/katsun0921/portfolio_api/src/domain/apis"
  "github.com/katsun0921/portfolio_api/src/services"
  "net/http"
)

func Get(c *gin.Context) {
	var resApi []*apis.Api
	var err rest_errors.RestErr

	webapi := c.Query("webapi")
	fmt.Println("query:",webapi)

	switch webapi {
    case "twitter":
      resApi, err = services.ApisService.GetTwitterApi()
    default:
      resApi, err = services.ApisService.GetApi()
  }

	if err != nil {
		logger.Error("error when trying to api request", err)
		restErr := rest_errors.NewBadRequestError("invalid json error", errors.New("json error"))
		if restErr != nil {
		  return
    }
	}
	c.JSON(http.StatusOK, resApi)
}
