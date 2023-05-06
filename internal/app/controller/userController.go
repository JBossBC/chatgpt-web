package controller

import (
	"chatgpt-web/internal/app/service"
	"time"

	"github.com/gin-gonic/gin"
)

type UserData struct {
	View
}

func LoginHandler(context *gin.Context) {
	// get
	username := context.Query("username")
	password := context.Query("password")
	if username == "" || password == "" {
		context.AbortWithStatusJSON(400, GetBadResponse("账号和密码为空"))
		return
	}

	result := service.NewUserService().Login(map[string]interface{}{"username": username, "password": password})
	context.AbortWithStatusJSON(200, GetSuccessResponse(map[string]interface{}{"data": result}))

}

func RegisterHandler(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")
	phone_number := context.Query("PhoneNumber")
	if username == "" || password == "" {
		context.AbortWithStatusJSON(400, GetBadResponse("账号和密码为空"))
		return
	}
	result := service.NewUserService().Register(map[string]interface{}{"username": username, "password": password, "PhoneNumber": phone_number, "Level": "nomal", "Create_time": time.Now(), "Update_time": time.Now()})
	context.AbortWithStatusJSON(200, GetSuccessResponse(map[string]interface{}{"data": result}))
}
