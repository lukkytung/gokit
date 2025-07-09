package redis

import (
	"log"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/lukkytung/gokit/pkg/config"
)

// Redis 客户端全局实例
var Client *redis.Client

// InitRedis 初始化 Redis 客户端
func InitRedis() {

	redisConfig := config.AppConfig
	addr := redisConfig.RedisHost + ":" + strconv.Itoa(redisConfig.RedisPort)
	password := redisConfig.RedisPassword
	db := redisConfig.RedisDb

	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// 测试 Redis 连接
	if _, err := Client.Ping().Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Redis connected successfully")
}
