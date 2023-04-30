package klog

import (
	"fmt"
	"io/fs"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"

	kafka "github.com/segmentio/kafka-go"
	"gopkg.in/yaml.v3"
)

type logConfig struct {
	Network       string     `yaml:Network`
	Address       string     `yaml:Address`
	StartConnPool bool       `yaml:StartConnPool`
	ConnsNumber   uint       `yaml:ConnsNumber`
	ConnsConfig   connConfig `yaml:ConnConfig`
}

type connConfig struct {
	ClientID        string
	Topic           string
	Partition       int32
	Broker          int32
	Rack            string
	TransactionalID string
}

const defaultConnsNumber = 100
const defaultStartConnPool = true
const defaultNetwork = "TCP"
const defaultEnvironment = "produce"

var config logConfig = logConfig{
	StartConnPool: defaultStartConnPool,
	ConnsNumber:   defaultConnsNumber,
	ConnsConfig:   connConfig{},
}

func init() {
	value := os.Getenv("chatgpt-web")
	if value == "" {
		value = defaultEnvironment
	}
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("log:log init error %s", err.Error()))
	}
	absPath := strings.Builder{}
	absPath.WriteString(dir)
	absPath.WriteRune(filepath.Separator)
	absPath.WriteString(value)
	absPath.WriteString(".yaml")
	file, err := os.OpenFile(absPath.String(), 644, fs.FileMode(os.O_RDONLY))
	if err != nil {
		panic(fmt.Sprintf("log:log init error %s", err.Error()))
	}
	if err := yaml.NewEncoder(file).Encode(&config); err != nil || config.Address == "" {
		panic(fmt.Sprintf("log:log init error %s", err.Error()))
	}
	initConnPool()
}
func initConnPool() {
	logCollect = &logCollection{
		dest:        make([]net.Conn, config.ConnsNumber),
		signal:      make(chan int, config.ConnsNumber),
		connsNumber: config.ConnsNumber,
		mu:          sync.Mutex{},
	}
	for i := 0; i < int(logCollect.connsNumber); i++ {
		dialer, err := net.Dial(config.Network, config.Address)
		if err != nil {
			panic(fmt.Sprintf("conns connection error: %s", err.Error()))
		}
		logCollect.dest[i] = kafka.NewConn(dialer, config.ConnsConfig.Topic, int(config.ConnsConfig.Partition))
		logCollect.signal <- i
	}
}

type logCollection struct {
	dest        []net.Conn
	signal      chan int
	mu          sync.Mutex
	connsNumber uint
}

func getIdleConn() int {
	return <-logCollect.signal
}

var logCollect *logCollection

func Fatal(info interface{}) {
	connIndex := getIdleConn()
	conn := logCollect.dest[connIndex]
	sb := strings.Builder{}
	sb.WriteString("Fatal Exception: ")
	sb.WriteString(fmt.Sprintf("%v", info))
	_, err := conn.Write([]byte(sb.String()))
	if err != nil {
		fmt.Printf("kafka client error:%v\n", err.Error())
	}
	logCollect.signal <- connIndex
}

func Error(info interface{}) {
	connIndex := getIdleConn()
	conn := logCollect.dest[connIndex]
	sb := strings.Builder{}
	sb.WriteString("Error Exception: ")
	sb.WriteString(fmt.Sprintf("%v", info))
	_, err := conn.Write([]byte(sb.String()))
	if err != nil {
		fmt.Printf("kafka client error:%v\n", err.Error())
	}
	logCollect.signal <- connIndex
}

func Print(info interface{}) {
	connIndex := getIdleConn()
	conn := logCollect.dest[connIndex]
	sb := strings.Builder{}
	sb.WriteString("Print info: ")
	sb.WriteString(fmt.Sprintf("%v", info))
	_, err := conn.Write([]byte(sb.String()))
	if err != nil {
		fmt.Printf("kafka client error:%v\n", err.Error())
	}
	logCollect.signal <- connIndex
}

func Warn(info interface{}) {
	connIndex := getIdleConn()
	conn := logCollect.dest[connIndex]
	sb := strings.Builder{}
	sb.WriteString("Warn: ")
	sb.WriteString(fmt.Sprintf("%v", info))
	_, err := conn.Write([]byte(sb.String()))
	if err != nil {
		fmt.Printf("kafka client error:%v\n", err.Error())
	}
	logCollect.signal <- connIndex
}
