package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var Red *redis.Client
var ctx = context.Background()

// 读取mysql连接的配置参数
func InitConfig() (host string, database string, port string, username string, password string) {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("读取配置文件错误", err)
	}
	host = viper.GetString("mysql.localhost")
	database = viper.GetString("mysql.database")
	port = viper.GetString("mysql.port")
	username = viper.GetString("mysql.username")
	password = viper.GetString("mysql.password")
	return host, database, port, username, password
}

// 初始化连接mysql
func InitMySQL() *gorm.DB {
	//自定义日志打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)
	host, database, port, username, password := InitConfig()
	urlList := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=true&loc=Local"
	fmt.Println("urlList===", urlList)
	DB, _ = gorm.Open(mysql.Open(urlList), &gorm.Config{Logger: newLogger})
	fmt.Println(" MySQL inited 。。。。")
	return DB

}

// 读取redis配置文件
func RedisConfig() (host string, selectDb int, port string, password string, polSize int, minIdleConn int) {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("读取redis配置文件错误", err)
	}
	host = viper.GetString("redis.host")
	selectDb = viper.GetInt("redis.selectDb")
	port = viper.GetString("redis.port")
	password = viper.GetString("redis.password")
	polSize = viper.GetInt("redis.polSize")
	minIdleConn = viper.GetInt("redis.minIdleConn")
	return host, selectDb, port, password, polSize, minIdleConn
}

// 初始化连接redis
func InitRedis() *redis.Client {
	host, selectDb, port, password, polSize, minIdleConn := RedisConfig()
	localhost := host + ":" + port
	Red = redis.NewClient(&redis.Options{
		Addr:         localhost,
		Password:     password,
		DB:           selectDb,
		PoolSize:     polSize,
		MinIdleConns: minIdleConn,
	})
	fmt.Println(" Redis inited 。。。。")
	return Red
}

// 设置redis
func SetRedis(key string, value string, t time.Duration) bool {
	t = t * 1000000000
	err := Red.Set(ctx, key, value, t).Err()
	fmt.Println("err数据打印", err)
	if err != nil {
		fmt.Println("错误打印", err)
		return false
	}
	fmt.Println("正确数据已缓存")
	return true
}

// 设置redis
func SetRedisInterface(key string, value *string, t time.Duration) bool {
	_, err := Red.Set(ctx, key, value, t).Result()
	fmt.Println("err数据打印", err)
	if err != nil {
		fmt.Println("错误打印", err)
		return false
	}
	fmt.Println("正确数据已缓存")
	return true
}

// 获取数据
func GetRedis(key string) string {
	result, err := Red.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return result
}

// 删除reids
func DelRedis(key string) bool {
	_, err := Red.Del(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// 延长过期时间
func ExpireRedis(key string, t int) bool {
	expire := time.Duration(t) * time.Second
	if err := Red.Expire(ctx, key, expire).Err(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
