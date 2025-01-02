package routes

import (
	"worklinker-api/controllers"

	"github.com/gin-gonic/gin"
)

func IndexRoute(router *gin.Engine) {
	indexGroup := router.Group("/index")

	{
		indexGroup.GET("/ping", controllers.Index)
	}
}
