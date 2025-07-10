package cmd

import (
	"log"

	"github.com/lukkytung/gokit/pkg/config"
	"github.com/lukkytung/gokit/pkg/service"
	"github.com/lukkytung/gokit/pkg/utils"
)

// InitGokit 初始化 Gokit
func InitGokit() {

	log.SetPrefix("[Gokit] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 初始化配置
	if err := config.InitConfig(); err != nil {
		log.Fatal("Error loading config")
	}

	// 初始化sonyflake
	utils.InitIDGenerator()

	// 初始化数据库连接
	if err := service.InitPostgres(); err != nil {
		log.Fatal("Error connecting to database")
	}

	// 初始化 Redis
	if err := service.InitRedis(); err != nil {
		log.Fatal("Error connecting to Redis")
	}
}
