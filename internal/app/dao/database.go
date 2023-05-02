package dao

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/context"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type config struct {
	mysql mysqlConfig `yaml:"Mysql"`
	redis redisConfig `yaml:"redis"`
}

type mysqlConfig struct {
	url string `yaml:"Url"`
}
type redisConfig struct {
	address  string `yaml:"Address"`
	database int    `yaml:"Database"`
	password string `yaml:Password`
}

var databaseConfig = &config{}

var gormDB *gorm.DB

var redisClient *redis.Client

func init() {
	sb := strings.Builder{}
	sb.WriteString("configs")
	sb.WriteRune(filepath.Separator)
	//TODO command line params to set
	sb.WriteString("log_produce.yaml")
	file, err := os.OpenFile(sb.String(), os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	err = yaml.NewDecoder(file).Decode(databaseConfig)
	if err != nil {
		panic(fmt.Sprintf("database init error:%s", err.Error()))
	}
	gormDB, err = gorm.Open(mysql.Open(databaseConfig.mysql.url), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("mysql connection init error:%s", err.Error()))
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     databaseConfig.redis.address,
		Password: databaseConfig.redis.password,
		DB:       databaseConfig.redis.database,
	})
	err = redisClient.Ping(context.Background()).Err()
	if err != nil {
		panic(fmt.Sprintf("redis ping error:%s", err.Error()))
	}
}

func NewMysqlConn() *gorm.DB {
	return gormDB
}

func NewRedisConn()
