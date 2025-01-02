package controllers

import (
	"net/http"
	"worklinker-api/helpers"
	"worklinker-api/services"

	"github.com/gin-gonic/gin"
)

func GetProfileHandler(c *gin.Context) {
	// Extract user ID from the request parameters
	userID := c.Param("id")
	if userID == "" {
		helpers.RespondWithError(c, http.StatusBadRequest, "User ID is required")
		return
	}

	// Call the GetProfile service
	user, err := services.GetProfile(userID)
	if err != nil {
		helpers.RespondWithError(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User profile retrieved successfully",
		"data":    user,
	})

}
