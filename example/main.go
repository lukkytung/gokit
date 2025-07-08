package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukkytung/gokit/example/router"
	"github.com/lukkytung/gokit/pkg/config"
	"github.com/lukkytung/gokit/pkg/database"
	"github.com/lukkytung/gokit/pkg/redis"
)

func main() {

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

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	router.InitRouter(r)

	r.Run(":" + config.AppConfig.ServerPort)

}
