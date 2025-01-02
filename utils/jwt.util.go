package utils

import (
	"time"
	"worklinker-api/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string) string {
	claims := jwt.MapClaims{
		"Email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, _ := token.SignedString([]byte(config.Config("JWT_SECRET_KEY")))
	return t
}
