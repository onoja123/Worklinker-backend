package config

import (
	"log"

	"os"
	"worklinker-api/types"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var AppConfig types.Config
var db *mongo.Database
var userCollection *mongo.Collection

func InitializeMongoDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using default environment variables")
	}

	// Initialize configuration
	AppConfig = types.Config{
		Port:      os.Getenv("PORT"),
		MongoURI:  os.Getenv("MONGO_URI"),
		Database:  os.Getenv("DB_NAME"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}

func GetDB() *mongo.Database {
	return db
}

func GetUserCollection() *mongo.Collection {
	return userCollection
}
