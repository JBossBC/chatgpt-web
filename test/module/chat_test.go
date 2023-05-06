package test

import (
	"chatgpt-web/internal/app/dao"
	"chatgpt-web/internal/app/service"
	"fmt"
	"testing"
)

func TestChat(t *testing.T) {
	info := service.NewChatService().Chat("xiyang", dao.Message{
		Role:    "user",
		Content: "请详细说一下golang这门语言",
		Name:    "none",
	})
	fmt.Println(info)
}
