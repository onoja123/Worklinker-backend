package services

import (
	"context"
	"errors"
	"log"
	"time"
	"worklinker-api/config"
	"worklinker-api/models"
	"worklinker-api/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitAuthService(collection *mongo.Collection) {
	userCollection = collection
}

func Signup(user *models.User) (string, error) {
	// Check if user already exists
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser models.User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return "", errors.New("user already exists with this email")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	// Insert user into database
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return "", errors.New("failed to create user")
	}

	// Generate JWT token
	token := utils.GenerateToken(user.Email)

	return token, nil
}

func Login(email, password string) (string, error) {
	// Fetch the user from the database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := config.GetUserCollection().FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Check if the password is correct
	if err := utils.CheckPasswordHash(user.Password, password); err != nil {
		return "", errors.New("invalid password")
	}

	// Generate JWT token
	token := utils.GenerateToken(user.Email)

	return token, nil
}
