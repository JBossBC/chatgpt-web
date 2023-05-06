package klog

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"

	kafka "github.com/segmentio/kafka-go"
	"gopkg.in/yaml.v3"
)

type LogConfig struct {
	Network          string     `yaml:"Network"`
	Address          string     `yaml:"Address"`
	StartConnPool    bool       `yaml:"StartConnPool"`
	ConnsNumber      uint32     `yaml:"ConnsNumber"`
	ConnsConfig      connConfig `yaml:"ConnConfig"`
	BufioStripNumber uint32     `yaml:"BufioStripNumber"`
}

type connConfig struct {
	ClientID        string `yaml:"ClientID"`
	Topic           string `yaml:"Topic"`
	Partition       int32  `yaml:"Partition"`
	Broker          int32  `yaml:"Broker"`
	Rack            string `yaml:"Rack"`
	TransactionalID string `yaml:"TransactionalID"`
}

const defaultConnsNumber = 100
const defaultStartConnPool = true
const defaultNetwork = "tcp"
const defaultEnvironment = "log_produce"

// const defaultStripNumber = 5

var config LogConfig = LogConfig{
	StartConnPool: defaultStartConnPool,
	ConnsNumber:   defaultConnsNumber,
	ConnsConfig:   connConfig{},
	// BufioStripNumber: defaultStripNumber,
	Network: defaultNetwork,
}

//TODO envrionment difference
func init() {
	value := os.Getenv("chatgpt-web")
	if value == "" {
		value = defaultEnvironment
	}
	// dir, err := os.Getwd()
	// if err != nil {
	// 	panic(fmt.Sprintf("log:log init error %s", err.Error()))
	// }
	absPath := strings.Builder{}
	// absPath.WriteString(dir)
	absPath.WriteString("./configs")
	absPath.WriteRune(filepath.Separator)
	absPath.WriteString(value)
	absPath.WriteString(".yaml")
	file, err := os.OpenFile(absPath.String(), os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("log:log init error %s", err.Error()))
	}
	if err := yaml.NewDecoder(file).Decode(&config); err != nil || config.Address == "" {
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
		drained:     make(chan struct{}, 1),
		// bufio:           bufio.Writer{},
		// bufioStipNumber: config.BufioStripNumber,
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
	connsNumber uint32
	drained     chan struct{}
	// bufio           bufio.Writer
	// bufioStipNumber uint32
	// counter         uint32
	// last  bit: reset State
	//penultimate bit: full state
	// writeState uint32
}

func getIdleConn() int {
	return <-logCollect.signal
}

var logCollect *logCollection

//TODO rollup and reset state
// func rollup(info interface{}) bool {
// 	logCollect.mu.Lock()
// 	defer logCollect.mu.Unlock()

// 	return false
// }
func Fatal(info interface{}) {
	go func() {
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
	}()
}

func Error(info interface{}) {
	go func() {
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
	}()
}

func Print(info interface{}) {
	go func() {
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
	}()
}

func Warn(info interface{}) {
	go func() {
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
	}()
}
