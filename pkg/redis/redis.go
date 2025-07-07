package redis

import (
	"context"
	"fmt"

	"github.com/lukkytung/gokit/pkg/config"
	"github.com/lukkytung/gokit/pkg/logger"

	"github.com/go-redis/redis/v8"
)

// Redis 客户端全局实例
var Client *redis.Client

// InitRedis 初始化 Redis 客户端
func InitRedis() {

	redisConfig := config.AppConfig
	addr := redisConfig.RedisHost + ":" + fmt.Sprint(redisConfig.RedisPort)
	password := redisConfig.RedisPassword
	db := redisConfig.RedisDb

	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// 测试 Redis 连接
	if _, err := Client.Ping(context.Background()).Result(); err != nil {
		logger.Log.Fatalf("Failed to connect to Redis: %v", err)
	}
	logger.Log.Info("Redis connected successfully")
}
