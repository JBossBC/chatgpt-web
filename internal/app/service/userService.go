package service

import (
	"chatgpt-web/internal/app/dao"
	"sync"
)

var (
	user *UserService
	once sync.Once = sync.Once{}
)

type UserService struct {
}

func NewUserService() *UserService {
	once.Do(func() {
		user = &UserService{}
	})
	return user
}

func (*UserService) Register() {

}

func (*UserService) Login(user dao.User) (info string) {
	defer func() {
		if err := recover(); err != nil {
			info = "系统错误"
		}
	}()
	dbUser, err := dao.QueryUser(user.Username)
	if err != nil {
		info = err.Error()
	}
	return info
}
func (*UserService) Register(user dao.User)
