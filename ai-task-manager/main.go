package main

import (
	"ai-task-manager/config"
	"ai-task-manager/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	r.POST("/register", routes.Register)
	r.POST("/login", routes.Login)

	r.Run(":8080")
}