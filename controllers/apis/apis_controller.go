package apis

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/katsun0921/go_utils/logger"
	"github.com/katsun0921/go_utils/rest_errors"
	"github.com/katsun0921/portfolio_api/domain/apis"
	"github.com/katsun0921/portfolio_api/services"
	"net/http"
)

func Get(c *gin.Context) {
	var api apis.Api

	rss, err := services.ApisService.GetApi(api)
	if err != nil {
		logger.Error("error when trying to api request", err)
		restErr := rest_errors.NewBadRequestError("invalid json error", errors.New("json error"))
		if restErr != nil {
		  return
    }
	}
	c.JSON(http.StatusOK, rss)
}
