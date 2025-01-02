package helpers

import "github.com/gin-gonic/gin"

func RespondWithError(c *gin.Context, status int, message string) error {
	c.JSON(status, gin.H{
		"status":  status,
		"success": false,
		"error":   message,
	})
	return nil
}

func RespondWithSuccess(c *gin.Context, status int, message string) error {
	c.JSON(status, gin.H{
		"status":  status,
		"success": true,
		"data":    message,
	})
	return nil
}
