package test

import (
	"chatgpt-web/internal/app/dao"
	"fmt"
	"testing"
	"time"
)

func TestUserInsert(t *testing.T) {
	fmt.Println(dao.NewUserDao().InsertUser(dao.User{
		Username:      "longshao",
		Password:      "123456",
		PhoneNumber:   "18080705675",
		Level:         "normal",
		LevelDeadline: time.Time{},
	}))
}
func TestUserQuery(t *testing.T) {
	fmt.Println(dao.NewUserDao().QueryUser("longshao"))
}

func TestUserUpdate(t *testing.T) {
	fmt.Println(dao.NewUserDao().UpdateUser(dao.User{
		Username: "test",
		Password: "3333444",
	}))
}
func TestUserDelete(t *testing.T) {
	fmt.Println(dao.NewUserDao().DeleteUser(dao.User{
		Username: "longshao",
		Password: "3333444",
	}))
	time.Sleep(5 * time.Second)
}
