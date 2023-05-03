package dao

import (
	klog "chatgpt-web/internal/app/log"
	"chatgpt-web/internal/app/task"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

//TODO password encrypto
type User struct {
	Id            uint64    `gorm:"column:id;primaryKey;<-:false"`
	Username      string    `gorm:"column:username;index:"`
	Password      string    `gorm:"column:password"`
	PhoneNumber   string    `gorm:"column:phoneNumber"`
	Level         string    `gorm:"column:Level"`
	LevelDeadline time.Time `gorm:"column:levelDeadline"`
	Create_time   time.Time `gorm:"column:createTime"`
	Update_time   time.Time `gorm:"column:updateTime"`
	Delete_time   time.Time `gorm:"column:deleteTime"`
}

func init() {
	newMysqlConn().AutoMigrate(&User{})
}

func NewUser(username string) User {
	return User{
		Username: username,
	}
}
func (user *User) WithLevel(level string) {
	user.Level = level
}
func (user *User) WithPhoneNumber(phoneNumber string) {
	user.PhoneNumber = phoneNumber
}
func (user *User) WithLevelDeadline(time time.Time) {
	user.LevelDeadline = time
}
func (user *User) WithPassword(password string) {
	user.Password = password
}
func (user *User) WithCreateTime(time time.Time) {
	user.Create_time = time
}
func (user *User) WithUpdateTime(time time.Time) {
	user.Update_time = time
}
func (user *User) WithDeleteTime(time time.Time) {
	user.Delete_time = time
}

type userDao struct {
}

var (
	user *userDao
	once sync.Once
)

func NewUserDao() *userDao {
	once.Do(func() {
		user = &userDao{}
	})
	return user
}

func (*userDao) InsertUser(user User) error {
	var existUser User
	err := newMysqlConn().Select("username").Where("username=? and deleteTime is NULL ", user.Username).First(&existUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("系统错误，请稍后重试")
	}
	var exists = existUser != User{}
	if !exists {
		err = newMysqlConn().Create(&user).Error
		if err != nil {
			klog.Error(err)
			return fmt.Errorf("系统错误，请稍后重试")
		}
		klog.Print(fmt.Sprintf("创建用户成功: %v", user))
		ctx, cancel := context.WithTimeout(context.Background(), defaultQueryTimeout)
		defer cancel()
		jsonRes, err := json.Marshal(&user)
		if err != nil {
			klog.Error(err)
			return nil
		}
		newRedisConn().SetEx(ctx, user.Username, jsonRes, defaultKeyTimeout)
	} else {
		return fmt.Errorf("用户名已存在")
	}
	return nil
}
func (*userDao) UpdateUser(user User) error {
	err := newMysqlConn().Where("username=?", user.Username).Updates(user).Error
	if err != nil {
		klog.Error(err)
		return fmt.Errorf("系统错误，请稍后重试")
	}
	klog.Print(fmt.Sprintf("修改用户成功:%v", user))
	var retryTask = func() error {
		_, err = newRedisConn().Del(context.Background(), user.Username).Result()
		//TODO defer query task
		if err != nil && err != redis.Nil {
			klog.Error(err)
			return err
		}
		return nil
	}
	task.AddTask(retryTask, task.HighPriority)
	return nil
}
func (*userDao) DeleteUser(user User) error {
	deleteUser := User{
		Username:    user.Username,
		Delete_time: time.Now(),
	}
	err := newMysqlConn().Where("username=?", user.Username).Updates(deleteUser).Error
	if err != nil {
		klog.Error(err)
		return fmt.Errorf("系统错误，请稍后重试")
	}
	klog.Print(fmt.Sprintf("软删除用户成功: %v", user))
	var retryTask = func() error {
		_, err = newRedisConn().Del(context.Background(), user.Username).Result()
		//TODO defer query task
		if err != nil && err != redis.Nil {
			return err
		}
		return nil
	}
	task.AddTask(retryTask, task.HighPriority)
	return nil
}
func (*userDao) QueryUserList(username []string) ([]User, error) {
	return nil, nil
}

const defaultInsertTimeout = 500 * time.Millisecond
const defaultQueryTimeout = 500 * time.Millisecond
const defaultKeyTimeout = 60 * 5 * time.Second
const defaultInvalidKeyTimeout = 60 * time.Second

func (*userDao) QueryUser(username string) (user User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultQueryTimeout)
	defer cancel()
	result, err := newRedisConn().Get(ctx, username).Result()
	if result == "" {
		err := newMysqlConn().Where("username=? and deleteTime is NULL ", username).First(&user).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			klog.Error(err)
			return user, fmt.Errorf("系统繁忙,请稍后重试")
		}
		isZero := user == User{}
		if !isZero {
			jsonRes, err := json.Marshal(user)
			if err != nil {
				klog.Error(err)
				return user, nil
			}
			newRedisConn().SetEx(context.Background(), username, jsonRes, defaultKeyTimeout)
			return user, nil
		} else {
			//cache
			newRedisConn().SetEx(context.Background(), username, InvalidKeyValue, defaultInvalidKeyTimeout)
			return user, fmt.Errorf("用户名不存在")
		}
	} else if result != InvalidKeyValue {
		err = json.Unmarshal([]byte(result), &user)
		if err != nil {
			klog.Error(fmt.Sprintf("json unmarshal error: %s(%s)", err.Error(), result))
		}
		return user, nil
	} else {
		return user, fmt.Errorf("用户名不存在")
	}
}
