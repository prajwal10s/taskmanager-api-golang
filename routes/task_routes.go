package routes

import (
	"taskmanager/controllers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	taskGroup := r.Group("/tasks")
	{
		taskGroup.POST("/", controllers.CreateTask)        // To create a task
		taskGroup.GET("/", controllers.GetTasks)           // To list tasks with optional query parameters
		taskGroup.GET("/:id", controllers.GetTaskByID)     // To get a specific task by ID
		taskGroup.PATCH("/:id", controllers.UpdateTask)    // To update a task
		taskGroup.DELETE("/:id", controllers.DeleteTask)   // To delete a task
	}
}