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

// 此处设计得比较差，忘记了处理返回得状态该怎么交予用户,仅仅只是返回了一个info,入参也对context不可见
func (*UserService) Login(data map[string]interface{}) (info string, cookie []byte) {
	defer func() {
		if err := recover(); err != nil {
			klog.Fatal(err)
			info = "系统错误"
		}
	}()
	value, err := utils.Marshal(data, reflect.TypeOf(dao.User{}))
	if err != nil {
		klog.Error(err)
		return "系统错误", []byte("")
	}
	user := value.Interface().(*dao.User)
	dbUser, err := dao.NewUserDao().QueryUser(user.Username)
	if err != nil {
		info = err.Error()
	}
	if utils.Encryption(user.Password) != dbUser.Password {
		return "登录失败", []byte("")
	} else {
		jwt, err := utils.MarshalJWT(dbUser)
		if err != nil {
			return "登录失败", []byte("")
		}
		return "登录成功", jwt
	}
}
func (*UserService) Register(data map[string]interface{}) (info string, cookie []byte) {
	defer func() {
		if err := recover(); err != nil {
			klog.Fatal(err)
			info = "系统错误"
		}
	}()
	value, err := utils.Marshal(data, reflect.TypeOf(dao.User{}))
	if err != nil {
		klog.Error(err)
		return "系统错误", []byte("")
	}
	user := *value.Interface().(*dao.User)
	user.Password = utils.Encryption(user.Password)
	nowTime := time.Now()
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "系统正在开小差,请稍后", []byte("")
	}
	nowTime = nowTime.In(loc)
	user.Create_time = nowTime
	user.Update_time = nowTime
	err = dao.NewUserDao().InsertUser(user)
	if err != nil {
		return err.Error(), []byte("")
	}
	jwt, _ := utils.MarshalJWT(user)
	return "注册成功", jwt
}
