package service

import (
	"chatgpt-web/internal/app/dao"
	klog "chatgpt-web/internal/app/log"
	"chatgpt-web/internal/app/utils"
	"reflect"
	"sync"
	"time"
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

func (*UserService) Login(data map[string]interface{}) (info string) {
	defer func() {
		if err := recover(); err != nil {
			klog.Fatal(err)
			info = "系统错误"
		}
	}()
	value, err := utils.Marshal(data, reflect.TypeOf(dao.User{}))
	if err != nil {
		klog.Error(err)
		return "系统错误"
	}
	user := value.Interface().(*dao.User)
	dbUser, err := dao.NewUserDao().QueryUser(user.Username)
	if err != nil {
		info = err.Error()
	}
	if utils.Encryption(user.Password) != dbUser.Password {
		return "登录失败"
	} else {
		return "登录成功"
	}
}
func (*UserService) Register(data map[string]interface{}) (info string) {
	defer func() {
		if err := recover(); err != nil {
			klog.Fatal(err)
			info = "系统错误"
		}
	}()
	value, err := utils.Marshal(data, reflect.TypeOf(dao.User{}))
	if err != nil {
		klog.Error(err)
		return "系统错误"
	}
	user := *value.Interface().(*dao.User)
	user.Password = utils.Encryption(user.Password)
	nowTime := time.Now()
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "系统正在开小差,请稍后"
	}
	nowTime = nowTime.In(loc)
	user.Create_time = nowTime
	user.Update_time = nowTime
	err = dao.NewUserDao().InsertUser(user)
	if err != nil {
		return err.Error()
	}
	return "注册成功"
}
