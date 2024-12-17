package controllers

import (
	"context"
	"net/http"
	"time"

	"taskmanager/db"
	"taskmanager/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateTask - Create a task
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(models.User)
	task.ID = primitive.NewObjectID()
	task.Owner = user.ID

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.DB.Collection("tasks").InsertOne(ctx, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// GetTasks - List tasks
func GetTasks(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := db.DB.Collection("tasks").Find(ctx, bson.M{"owner": user.ID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	var tasks []models.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID - Get a specific task
func GetTaskByID(c *gin.Context) {
	// Similar to above
}

// UpdateTask - Update a task
func UpdateTask(c *gin.Context) {
	// Similar to user update logic
}

// DeleteTask - Delete a task
func DeleteTask(c *gin.Context) {
	// Similar to user delete logic
}