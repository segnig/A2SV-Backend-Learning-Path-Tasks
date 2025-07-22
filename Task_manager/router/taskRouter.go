package router

import (
	"task-manager/controllers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(incomingRoutes *gin.Engine) {
	api := incomingRoutes.Group("/api")
	{
		api.GET("/tasks", controllers.GetTasks)
		api.GET("/tasks/:id", controllers.GetTaskById)
		api.POST("/tasks", controllers.CreateTask)
		api.PUT("/tasks/:id", controllers.UpdateTask)
		api.DELETE("tasks/:id", controllers.DeleTeTask)
	}
}
