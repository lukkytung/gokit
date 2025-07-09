package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukkytung/gokit/example/model"
	"github.com/lukkytung/gokit/example/router"
	"github.com/lukkytung/gokit/pkg/cmd"
	"github.com/lukkytung/gokit/pkg/config"
	"github.com/lukkytung/gokit/pkg/database"
)

func main() {
	// 初始化Gokit
	cmd.Init()

	// 自动迁移数据库
	database.DB.AutoMigrate(&model.User{})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gokit!"})
	})

	router.InitRouter(r)

	r.Run(":" + config.AppConfig.ServerPort)

}
