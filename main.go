package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	//importing dotenv to use environment variables
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	//check for error while loading the file
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	// Set up our GIN router
	r := gin.Default()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server running on port %s", port)
	err = r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}