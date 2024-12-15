package routes

import (
	"taskmanager/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", controllers.RegisterUser) // To create a new user
		userGroup.POST("/login", controllers.LoginUser) // To log in a user
		userGroup.GET("/me", controllers.GetUser) // To fetch user 
		userGroup.PATCH("/me", controllers.UpdateUser) // To update user 
		userGroup.DELETE("/me", controllers.DeleteUser) // To delete user 
	}
}