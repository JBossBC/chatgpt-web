package test

import (
	"chatgpt-web/internal/app/dao"
	"fmt"
	"testing"
)

func TestUserInsert(t *testing.T) {
	fmt.Println(dao.NewUserDao().InsertUser(dao.User{
		Username:      "test",
		Password:      "123456",
		PhoneNumber:   "18080705675",
		Level:         "normal",
		LevelDeadline: nil,
	}))
}
func TestUserQuery(t *testing.T) {
	fmt.Println(dao.NewUserDao().QueryUser("test"))
}

func TestUserUpdate(t *testing.T) {
	fmt.Println(dao.NewUserDao().UpdateUser(dao.User{
		Username: "test",
		Password: "test",
	}))
}
