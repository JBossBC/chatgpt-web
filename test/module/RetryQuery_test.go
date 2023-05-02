package test

import (
	"chatgpt-web/internal/app/task"
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
			}, int8(i%256))
		}
	}()
	go task.Run()
	time.Sleep(30 * time.Second)
}
