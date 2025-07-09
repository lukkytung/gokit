package cmd

import (
	"log"

	"github.com/lukkytung/gokit/pkg/config"
	"github.com/lukkytung/gokit/pkg/database"
	"github.com/lukkytung/gokit/pkg/redis"
	"github.com/lukkytung/gokit/pkg/utils"
)

func InitGokit() {

	// 初始化sonyflake
	utils.InitIDGenerator()

	// 初始化配置
	if err := config.InitConfig(); err != nil {
		log.Fatal("Error loading config")
	}

	// 初始化数据库连接
	if err := database.InitPostgres(); err != nil {
		log.Fatal("Error connecting to database")
	}

	// 初始化 Redis
	redis.InitRedis()
}
