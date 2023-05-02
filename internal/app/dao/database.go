package dao

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type config struct {
	Mysql mysqlConfig `yaml:"Mysql"`
	Redis redisConfig `yaml:"Redis"`
}

type mysqlConfig struct {
	Url string `yaml:"Url"`
}
type redisConfig struct {
	Address  string `yaml:"Address"`
	Database int    `yaml:"Database"`
	Password string `yaml:"Password"`
}

var databaseConfig = &config{}

var gormDB *gorm.DB

var redisClient *redis.Client

//TODO envrionment difference
func init() {
	sb := strings.Builder{}
	// wd, err := os.Getwd()
	sb.WriteString("D:\\chatgpt-web")
	sb.WriteRune(filepath.Separator)
	sb.WriteString("configs")
	sb.WriteRune(filepath.Separator)
	//TODO command line params to set
	sb.WriteString("database_produce.yaml")
	file, err := os.OpenFile(sb.String(), os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	err = yaml.NewDecoder(file).Decode(databaseConfig)
	if err != nil {
		panic(fmt.Sprintf("database init error:%s", err.Error()))
	}
	gormDB, err = gorm.Open(mysql.Open(databaseConfig.Mysql.Url), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("mysql connection %s init error:%s ", databaseConfig.Mysql.Url, err.Error()))
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     databaseConfig.Redis.Address,
		Password: databaseConfig.Redis.Password,
		DB:       databaseConfig.Redis.Database,
	})
	err = redisClient.Ping(context.Background()).Err()
	if err != nil {
		panic(fmt.Sprintf("redis ping error:%s", err.Error()))
	}
}

func newMysqlConn() *gorm.DB {
	return gormDB
}

func newRedisConn() *redis.Conn {
	return redisClient.Conn()
}
