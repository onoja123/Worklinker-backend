package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"worklinker-api/config"
	"worklinker-api/helpers"
	"worklinker-api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helpers.RespondWithError(c, http.StatusUnauthorized, "Missing Authorization")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			helpers.RespondWithError(c, http.StatusUnauthorized, "Invalid token format")
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Config("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			helpers.RespondWithError(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := claims["userId"].(string)
			parsedID, _ := uuid.Parse(userId)
			err = userCollection.FindOne(context.Background(), bson.M{"_id": parsedID}).Decode(&user)
			if err == nil {
				c.Set("user", user)
			} else {
				helpers.RespondWithError(c, http.StatusUnauthorized, "User not found")
				c.Abort()
				return
			}
		} else {
			helpers.RespondWithError(c, http.StatusUnauthorized, "Invalid token claims")
			c.Abort()
			return
		}

		c.Next()
	}
}
