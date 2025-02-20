package main

import (
	"ai-task-manager/config"
	"ai-task-manager/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	// Authentication
	r.POST("/register", routes.Register)
	r.POST("/login", routes.Login)

	// Tasks
	r.GET("/tasks", routes.GetTasks)
	r.POST("/tasks", routes.CreateTask)
	r.GET("/ws", routes.WebSocketHandler)

	// Start WebSocket broadcaster
	go routes.TaskBroadcaster()

	r.Run(":8080")
}
