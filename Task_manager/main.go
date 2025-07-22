package main

import (
	"os"
	routes "task-manager/router"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.TaskRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)

}
