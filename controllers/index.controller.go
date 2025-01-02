package controllers

import (
	"net/http"
	"worklinker-api/helpers"
	"worklinker-api/services"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	response := services.Ping()
	helpers.RespondWithSuccess(
		c,
		http.StatusOK,
		response,
	)

}
