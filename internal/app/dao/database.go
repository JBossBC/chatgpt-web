package dao

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const InvalidKeyValue = "@invalid@Key@Value@"

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

const defaultDatabaseConfig = "./config/database.yaml"

//TODO envrionment difference
func init() {
	value := os.Getenv("chatgpt-web-database")
	if value == "" {
		value = defaultDatabaseConfig
	}
	// wd, err := os.Getwd()
	//TODO command line params to set
	fmt.Printf("initing the database config %s\n", value)
	file, err := os.OpenFile(value, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
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
