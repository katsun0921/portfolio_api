package app

import (
  "github.com/gin-gonic/gin"
  "github.com/katsun0921/portfolio_api/logger"
)

var(
  router = gin.Default()
)

func StartApplication() {
  mapUrls()

  logger.Info("about to start the application...")
  err := router.Run(":8081")
  if err != nil {
    return
  }
}
