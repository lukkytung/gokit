package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lukkytung/gokit/example/controller"
)

func InitRouter(r *gin.Engine) {
	r.POST("/send-code", controller.SendCode)
	r.POST("/login", controller.LoginWithCode)
}
