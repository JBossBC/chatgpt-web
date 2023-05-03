package test

import (
	"chatgpt-web/internal/app/task"
	"errors"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)

func TestRetryQuery(t *testing.T) {
	task := task.NewRetryQuery(30 * time.Millisecond)
	var defaultMaxTimes = 1000
	go func() {
		for i := 0; i < defaultMaxTimes; i++ {
			task.AddTask(func() error {
				println(i)
				return nil
			}, 0)
		}
	}()
	go task.Run()
	time.Sleep(30 * time.Second)
}

var taskObj = task.NewRetryQuery(30 * time.Millisecond)

func TestClosure(t *testing.T) {
	var defaultMaxTimes = 1000
	for i := 0; i < defaultMaxTimes; i++ {
		closure(strconv.FormatInt(int64(i), 10))
	}
	go taskObj.Run()
	time.Sleep(30 * time.Second)
	println(times)
}

var times int64 = 0
var errorTimes int64 = 1

func closure(user string) {
	var retry = func() error {
		atomic.AddInt64(&times, 1)
		if errorTimes < 5 {
			atomic.AddInt64(&errorTimes, 1)
			return errors.New("hello")
		} else {
			errorTimes = 1
		}
		println(user)
		return nil
	}
	taskObj.AddTask(retry, 100)
}
