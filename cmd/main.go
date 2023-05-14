package main

import (
	"chatgpt-web/internal/app/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.Use()
	chatGroup := engine.Group("/chat")
	{
		chatGroup.POST("/", controller.ChatHandler)
	}
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register", controller.RegisterHandler)
		userGroup.GET("/login", controller.LoginHandler)
	}
	engine.Run(":8080")
}
