package controllers

import (
	"net/http"
	"worklinker-api/config"
	"worklinker-api/helpers"
	"worklinker-api/models"
	"worklinker-api/services"
	"worklinker-api/types"

	"github.com/gin-gonic/gin"
)

func SignupHandler(c *gin.Context) {
	var user models.User

	// Bind JSON to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate required fields
	if user.Email == "" || user.Password == "" {
		helpers.RespondWithError(c, http.StatusBadRequest, "Email and Password are required")
		return
	}

	// Call the signup service
	token, err := services.Signup(&user)
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "error")
		return
	}

	// Create a wallet for the user
	walletService := services.NewWalletService(config.GetDB())
	_, err = walletService.CreateWallet(user.ID, "USD")
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to create wallet")
		return
	}

	helpers.RespondWithSuccess(
		c,
		http.StatusOK,
		token,
	)

}

func LoginHandler(c *gin.Context) {

	var payload types.LoginPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate fields
	if payload.Email == "" || payload.Password == "" {
		helpers.RespondWithError(c, http.StatusBadRequest, "Email and Password are required")
		return
	}

	// Call the login service
	token, err := services.Login(payload.Email, payload.Password)
	if err != nil {
		helpers.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
