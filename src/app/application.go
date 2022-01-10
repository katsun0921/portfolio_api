package app

import (
  "github.com/gin-gonic/gin"
  "github.com/katsun0921/go_utils/logger"
)

var(
  router = gin.Default()
)

func StartApplication() {
  mapUrls()

  logger.Info("about to start the application...")
  router.Run(":3000")
}
