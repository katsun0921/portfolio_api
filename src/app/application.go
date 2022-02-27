package app

import (
	"github.com/gin-gonic/gin"
	"github.com/katsun0921/go_utils/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(CORSMiddleware())

	mapUrls()

	logger.Info("about to start the application...")
	router.Run(":3000")
}


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")


		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
