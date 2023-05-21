package main

import (
	"chatgpt-web/internal/app/controller"
	"chatgpt-web/internal/app/middlerware"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.Use(middlerware.JWT)
	chatGroup := engine.Group("/chat")
	{
		chatGroup.GET("/", controller.ChatHandler)
	}
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register", controller.RegisterHandler)
		userGroup.GET("/login", controller.LoginHandler)
	}
	engine.Run(":8080")
}
