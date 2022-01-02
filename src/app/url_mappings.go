package app

import (
	"github.com/katsun0921/portfolio_api/src/controllers/apis"
  "github.com/katsun0921/portfolio_api/src/controllers/articles"
  "github.com/katsun0921/portfolio_api/src/controllers/ping"
)

func mapUrls() {
  router.GET("/ping", ping.Ping)
  router.GET("/apis", apis.Get)
  router.GET("/articles", articles.Get)
  router.POST("/articles", articles.Create)
}
