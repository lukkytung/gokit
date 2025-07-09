package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukkytung/gokit/example/model"
	"github.com/lukkytung/gokit/example/router"
	"github.com/lukkytung/gokit/pkg/cmd"
	"github.com/lukkytung/gokit/pkg/config"
	"github.com/lukkytung/gokit/pkg/service"
)

func main() {
	// 初始化Gokit
	cmd.InitGokit()

	// 自动迁移数据库
	service.DB.AutoMigrate(&model.User{})

	// 初始化 Gin
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gokit!"})
	})

	router.InitRouter(r)

	// 加载 templates 目录下的 HTML 文件
	r.LoadHTMLGlob("templates/*")

	r.Run(":" + config.AppConfig.ServerPort)

}
