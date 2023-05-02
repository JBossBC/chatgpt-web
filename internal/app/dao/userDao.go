package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Id            uint64    `gorm:"column:id;primaryKey;<-:false"`
	Username      string    `gorm:"column:username"`
	Password      string    `gorm:"column:password"`
	PhoneNumber   string    `gorm:"column:phoneNumber"`
	Level         string    `gorm:"column:Level"`
	LevelDeadline time.Time `gorm:"column:levelDeadline"`
	Create_time   time.Time `gorm:"column:createTime"`
	Update_time   time.Time `gorm:"column:updateTime"`
	Delete_time   time.Time `gorm:"column:deleteTime"`
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

func InsertUser(user User) error {
	var existUser User
	err := NewMysqlConn().Select("username").Where("username=? and deleteTime=Null", user.Username).First(&existUser).Error
	if err != nil {
		return fmt.Errorf("系统错误，请稍后重试")
	}
	var exists = existUser != User{}
	if !exists {
		err = NewMysqlConn().Create(&user).Error
		if err != nil {
			return fmt.Errorf("系统错误，请稍后重试")
		}
		ctx, cancel := context.WithTimeout(context.Background(), defaultQueryTimeout)
		defer cancel()
		jsonRes, err := json.Marshal(&user)
		if err != nil {
			return nil
		}
		NewRedisConn().SetEx(ctx, user.Username, jsonRes, 10*time.Second)
	} else {
		return fmt.Errorf("用户名已存在")
	}
	return nil
}
func UpdateUser(user User) error {
	err := NewMysqlConn().Where("username=?", user.Username).Updates(user).Error
	if err != nil {
		return fmt.Errorf("系统错误，请稍后重试")
	}
	val, err := NewRedisConn().Del(context.Background(), user.Username).Result()
	if err == nil {
		var delay = func() {
			val, err := NewRedisConn().Del(context.Background(), user.Username).Result()
		}
	}
}
func DeleteUser(user User) error {

}
func QueryUserList(username []string) ([]User, error) {
}

const defaultInsertTimeout = 500 * time.Millisecond
const defaultQueryTimeout = 500 * time.Millisecond

func QueryUser(username string) (user User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultQueryTimeout)
	defer cancel()
	result, _ := NewRedisConn().Get(ctx, username).Result()
	if result == "" {
		err := NewMysqlConn().Where("username=? and deleteTime=Null ", username).First(&user).Error
		if err != nil {
			return user, fmt.Errorf("系统繁忙,请稍后重试")
		}
		isZero := user == User{}
		if !isZero {
			jsonRes, err := json.Marshal(user)
			if err != nil {
				return user, err
			}
			NewRedisConn().SetEx(context.Background(), username, jsonRes, 10*time.Second)
		}
		return user, err
	} else {
		err = json.Unmarshal([]byte(result), &user)
	}
	return user, err
}
