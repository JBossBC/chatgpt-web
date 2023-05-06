package service

import (
	"chatgpt-web/internal/app/dao"
	klog "chatgpt-web/internal/app/log"
	"fmt"
	"sync"
)

type chatService struct {
}

var (
	chat     *chatService
	chatOnce sync.Once
)

func NewChatService() *chatService {
	chatOnce.Do(func() {
		chat = &chatService{}
	})
	return chat
}

func (chat *chatService) Chat(key string, message dao.Message) (info string) {
	defer func() {
		if errPanic := recover(); errPanic != nil {
			klog.Fatal(errPanic)
			info = "系统繁忙"
		}
	}()
	result, err := dao.GetChatgptClient().SendQuestion(key, message)
	if err != nil {
		klog.Error(fmt.Sprintf("chatServer errror: %s", err.Error()))
		return "chatgpt正在开小差，请稍等"
	}
	return result
}
