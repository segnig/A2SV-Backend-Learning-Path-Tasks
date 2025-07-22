package router

import (
	"task-manager/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	api := incomingRoutes.Group("/api")
	{
		api.POST("users/signup", controllers.Signup())
		api.POST("users/login", controllers.Login())
	}
}
