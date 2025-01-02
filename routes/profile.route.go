package routes

import (
	"worklinker-api/controllers"

	"github.com/gin-gonic/gin"
)

func ProfileRoute(router *gin.Engine) {
	ProfileGroup := router.Group("/profile")

	{
		ProfileGroup.GET("/profile", controllers.GetProfileHandler)
	}
}
