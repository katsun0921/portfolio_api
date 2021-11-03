package app

import (
  "github.com/katsun0921/portfolio_api/controllers/apis"
  "github.com/katsun0921/portfolio_api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
  router.GET("/apis", apis.Get)
}
