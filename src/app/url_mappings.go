package app

import (
	"github.com/katsun0921/portfolio_api/src/controllers/apis"
	"github.com/katsun0921/portfolio_api/src/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/blogs", apis.GetBlogs)
	router.GET("/skills", apis.GetSkills)
	router.GET("/workExpress", apis.GetWorkExpress)
	//:TODO Comment out until post is made.
	// router.GET("/articles", articles.Get)
	// router.POST("/articles", articles.Create)
}
