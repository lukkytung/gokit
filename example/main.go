package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukkytung/gokit/example/model"
	"github.com/lukkytung/gokit/example/router"
	gokit "github.com/lukkytung/gokit/pkg/cmd"
	"github.com/lukkytung/gokit/pkg/config"
)

func main() {
	// 初始化Gokit
	gokit.InitWithModel(&model.User{})

	// 初始化 Gin
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gokit!"})
	})

	router.InitRouter(r)

	// 加载 templates 目录下的 HTML 文件
	r.LoadHTMLGlob("templates/*")

	r.Run(":" + config.AppConfig.ServerPort)

}
