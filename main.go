package main

import (
	"log"
	"worklinker-api/config"
	"worklinker-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.InitializeMongoDB()

	router := gin.Default()

	// Define routes
	routes.IndexRoute(router)
	routes.AuthRoutes(router)

	// Start the server
	port := config.Config("PORT")
	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
