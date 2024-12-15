package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//lets setup our mongodb

var DB *mongo.Database

func Connect() error{
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
	return nil
}

// CreateIndexes creates unique indexes for the users and tasks collections
func CreateIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCollection := DB.Collection("users")
	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	if _, err := userCollection.Indexes().CreateOne(ctx, emailIndex); err != nil {
		log.Fatalf("Failed to create user index: %v", err)
	}
}

func CreateTaskIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	taskCollection := DB.Collection("tasks")
	descriptionIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "description", Value: 1}},
		Options: options.Index().SetUnique(false), // Change to true for unique descriptions
	}

	if _, err := taskCollection.Indexes().CreateOne(ctx, descriptionIndex); err != nil {
		log.Fatalf("Failed to create task index: %v", err)
	}
}