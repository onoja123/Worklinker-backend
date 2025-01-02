package routes

import (
	"worklinker-api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", controllers.SignupHandler)
		authGroup.POST("/login", controllers.LoginHandler)
	}
}
