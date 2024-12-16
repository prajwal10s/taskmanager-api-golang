package routes

import (
	"taskmanager/controllers"
	"taskmanager/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", controllers.RegisterUser) // To create a new user
		userGroup.POST("/login", controllers.LoginUser) // To log in a user
		userGroup.GET("/me", middleware.AuthMiddleware(), controllers.GetUser) // To fetch user 
		userGroup.PATCH("/me", middleware.AuthMiddleware(), controllers.UpdateUser) // To update user 
		userGroup.DELETE("/me", middleware.AuthMiddleware(), controllers.DeleteUser) // To delete user 
	}
}