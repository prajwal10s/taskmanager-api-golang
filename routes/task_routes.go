package routes

import (
	"taskmanager/controllers"
	"taskmanager/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	taskGroup := r.Group("/tasks")
	{
		taskGroup.POST("/", middleware.AuthMiddleware(), controllers.CreateTask)        // To create a task
		taskGroup.GET("/", middleware.AuthMiddleware(), controllers.GetTasks)           // To list tasks with optional query parameters
		taskGroup.GET("/:id", middleware.AuthMiddleware(), controllers.GetTaskByID)     // To get a specific task by ID
		taskGroup.PATCH("/:id", middleware.AuthMiddleware(), controllers.UpdateTask)    // To update a task
		taskGroup.DELETE("/:id", middleware.AuthMiddleware(), controllers.DeleteTask)   // To delete a task
	}
}