package controller

import (
	"chatgpt-web/internal/app/service"

	"github.com/gin-gonic/gin"
)

type UserData struct {
	View
}

func loginHandler(context *gin.Context) {
	// get
	username := context.Query("username")
	password := context.Query("password")
	if username == "" || password == "" {
		context.AbortWithStatusJSON(400, GetBadResponse(400, "账号和密码为空"))
		return
	}

	service.NewUserService().Login(map[string]interface{}{"username": username, "password": password})

}
