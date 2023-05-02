package test

import (
	klog "chatgpt-web/internal/app/log"
	"testing"
	"time"
)

func TestKafkaClient(t *testing.T) {
	klog.Fatal("hello world")
	time.Sleep(5 * time.Second)
}
