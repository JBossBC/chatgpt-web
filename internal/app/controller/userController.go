package controller

import (
	"chatgpt-web/internal/app/service"
	"chatgpt-web/internal/app/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type UserData struct {
	utils.View
}

func LoginHandler(context *gin.Context) {
	// get
	username := context.Query("username")
	password := context.Query("password")
	if username == "" || password == "" {
		context.AbortWithStatusJSON(400, utils.GetBadResponse("账号和密码为空"))
		return
	}

	result, auth := service.NewUserService().Login(map[string]interface{}{"username": username, "password": password})
	if auth != nil {
		context.Writer.Header().Add("Authorization", string(auth))
	}
	context.AbortWithStatusJSON(200, utils.GetSuccessResponse(map[string]interface{}{"data": result}))

}

func RegisterHandler(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")
	phone_number := context.Query("PhoneNumber")
	if username == "" || password == "" {
		context.AbortWithStatusJSON(400, utils.GetBadResponse("账号和密码为空"))
		return
	}
	result, jwt := service.NewUserService().Register(map[string]interface{}{"username": username, "password": password, "PhoneNumber": phone_number, "Level": "nomal", "Create_time": time.Now(), "Update_time": time.Now()})
	if jwt != nil {
		context.Writer.Header().Add("Authorization", string(jwt))
	}
	context.AbortWithStatusJSON(200, utils.GetSuccessResponse(map[string]interface{}{"data": result}))
}
