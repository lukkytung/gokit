package main

import (
	"fmt"
	"log"

	"github.com/lukkytung/gokit/pkg/config"
	"github.com/lukkytung/gokit/pkg/database"
	"github.com/lukkytung/gokit/pkg/logger"
	"github.com/lukkytung/gokit/pkg/redis"
)

func main() {
	// 初始化配置
	if err := config.InitConfig("config.yaml"); err != nil {
		log.Fatal("Error loading config")
	}

	// 初始化日志
	logger.InitLogger()
	defer logger.Shutdown()

	// 初始化数据库连接
	if err := database.InitPostgres(); err != nil {
		log.Fatal("Error connecting to database")
	}

	// 初始化 Redis
	redis.InitRedis()

	// 示例输出
	fmt.Println("Application started successfully")
}
