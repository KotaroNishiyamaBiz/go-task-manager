package main

import (
	"github.com/gin-gonic/gin"
	"go-task-manager/controller"
)

func main() {
	router := gin.Default()

	// API namespace
	v1 := router.Group("/api/v1")
	{
		v1.GET("/tasks", controller.TasksGET)
		v1.POST("/tasks", controller.TaskPOST)
		v1.PATCH("/tasks/:id", controller.TaskPATCH)
		v1.DELETE("/tasks/:id", controller.TaskDELETE)
	}

	router.GET("/", controller.IndexGET)
	router.Run(":8080")
}