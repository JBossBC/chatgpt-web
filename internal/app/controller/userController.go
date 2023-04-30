package controller

import (
	"chatgpt-web/internal/app/service"
	"chatgpt-web/internal/pkg/log/klog"

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
		klog.Info()
		context.AbortWithStatusJSON(400, GetBadResponse(400, "账号和密码为空"))
		return
	}
	service.NewUserService().Login()

}
