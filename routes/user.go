package routes

import (
	"gogetters/controllers"
	"gogetters/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUser)
		api.POST("/users", controllers.CreateUser)
		api.DELETE("/users/:id", middleware.AuthMiddleware(), controllers.DeleteUser)
	}
}
