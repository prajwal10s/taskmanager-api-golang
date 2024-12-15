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

//Below are the routes needed according to the req.
	// 1.	POST /tasks - Create a task.
	// 2.	GET /tasks - List tasks with optional query parameters.
	// 3.	GET /tasks/:id - Get a specific task.
	// 4.	PATCH /tasks/:id - Update a task.
	// 5.	DELETE /tasks/:id - Delete a task.
	// 6.	POST /users - Create a new user.
	// 7.	POST /users/login - Log in a user.
	// 8.	POST /users/logout - Log out a user.
	// 9.	GET /users/me - Fetch user profile.
	// 10.	PATCH /users/me - Update user profile.
	// 11.	DELETE /users/me - Delete user account.


	
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