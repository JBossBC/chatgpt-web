package dao

import (
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

}
func UpdateUser(user User) error {

}
func DeleteUser(user User) error {

}
func QueryUserList(username []string) ([]User, error) {

}

func QueryUser(username string) (User, error) {

}
