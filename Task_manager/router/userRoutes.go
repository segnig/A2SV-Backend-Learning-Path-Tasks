package router

import (
	"task-manager/controllers"
	"task-manager/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	api := incomingRoutes.Group("/api")
	{
		api.Use(middleware.Authentication())
		api.GET("/users", controllers.GetUsers())
		api.GET("/users/:user_id", controllers.GetUser())
	}
}
