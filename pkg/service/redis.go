package service

import (
	"log"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/lukkytung/gokit/pkg/config"
)

// Redis 客户端全局实例
var RedisClient *redis.Client

// InitRedis 初始化 Redis 客户端
func InitRedis() error {

	redisConfig := config.AppConfig
	addr := redisConfig.RedisHost + ":" + strconv.Itoa(redisConfig.RedisPort)
	password := redisConfig.RedisPassword
	db := redisConfig.RedisDb

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// 测试 Redis 连接
	if _, err := RedisClient.Ping().Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
		return err
	}
	log.Println("Redis connected successfully")
	return nil
}
