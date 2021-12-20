package apis

import (
  "errors"
  "github.com/gin-gonic/gin"
  "github.com/katsun0921/go_utils/logger"
  "github.com/katsun0921/go_utils/rest_errors"
  "github.com/katsun0921/portfolio_api/src/constants"
  "github.com/katsun0921/portfolio_api/src/domain/apis"
  "github.com/katsun0921/portfolio_api/src/lib"
  "net/http"
)

func Get(c *gin.Context) {
  var resApi []*apis.Api
  var err rest_errors.RestErr

  service := c.Query(constants.QueryService)

  rssServices := [...]string{constants.ZENN}

  isRss := false
  for _, rssService := range rssServices {
    if rssService == service {
      isRss = true
      break
    }
  }

  if service == constants.TWITTER {
    resApi, err = lib.ApisService.GetTwitter()
  } else if isRss {
    resApi, err = lib.ApisService.GetRss(service)
  } else {
    resApi, err = lib.ApisService.GetApiAll()
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
