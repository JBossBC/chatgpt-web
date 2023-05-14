package controller

import (
	"chatgpt-web/internal/app/dao"
	"chatgpt-web/internal/app/service"
	"chatgpt-web/internal/app/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type ChatData struct {
	utils.View
}

func ChatHandler(context *gin.Context) {
	//jwt verify
	//middlerware make sure this element must have
	level := context.Keys["Level"]

	levelDeadLine := context.Keys["levelDeadline"]
	username := context.Keys["username"]

	nowTime := time.Now()
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		context.AbortWithStatusJSON(400, utils.GetBadResponse("系统正在开小差，请稍后"))
		return
	}
	nowTime = nowTime.In(loc)
	deadLine, err := time.Parse("2006-01-02 15:04:05", levelDeadLine.(string))
	if err != nil {
		context.AbortWithStatusJSON(400, utils.GetBadResponse("系统正在开小差，请稍后"))
		return
	}
	if level != "member" || nowTime.After(deadLine) {
		context.AbortWithStatusJSON(400, utils.GetBadResponse("你不是会员,成为会员后才开启问答功能"))
		return
	}
	contentStr := context.Query("content")
	info := service.NewChatService().Chat(username.(string), dao.Message{
		Role:    "user",
		Content: contentStr,
	})
	context.AbortWithStatusJSON(200, utils.GetSuccessResponse(map[string]interface{}{"data": info}))
}
