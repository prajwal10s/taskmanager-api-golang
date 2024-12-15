package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//lets setup our mongodb

var DB *mongo.Database

func Connect() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// MongoDB connection
	client, err := mongo.Connect(
		nil,
		options.Client().ApplyURI(os.Getenv("MONGODB_URL")),
	)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Set DB to a global variable
	DB = client.Database("taskmanagerGo")
	log.Println("Connected to MongoDB")
}