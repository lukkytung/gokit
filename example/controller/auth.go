package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lukkytung/gokit/example/model"
	"github.com/lukkytung/gokit/pkg/database"
	"github.com/lukkytung/gokit/pkg/jwt"
	"github.com/lukkytung/gokit/pkg/redis"
	"github.com/lukkytung/gokit/pkg/utils"
)

type CustomResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SendEmailRequest struct {
	Email string `json:"email"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func SendCode(c *gin.Context) {
	var req SendEmailRequest
	// 解析请求体中的 email
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, CustomResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid email",
			Data:    nil,
		})
		return
	}

	// 生成 6 位随机验证码
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	// 将验证码存入 Redis，有效期 10 分钟
	err := redis.Client.Set("code:"+req.Email, code, time.Minute*10).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, CustomResponse{
			Code:    http.StatusInternalServerError,
			Message: "Redis error",
			Data:    err.Error(),
		})
		return
	}

	// 发送邮件（忽略发送失败）
	err = utils.SendEmail(req.Email, "Your Login Code", "Your code is: "+code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CustomResponse{
			Code:    http.StatusInternalServerError,
			Message: "Email sending failed",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, CustomResponse{
		Code:    http.StatusOK,
		Message: "Code sent",
		Data:    gin.H{"code": code},
	})
}

func LoginWithCode(c *gin.Context) {

	db := database.DB

	var req LoginRequest
	// 解析请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, CustomResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request",
			Data:    nil,
		})
		return
	}

	// 校验 Redis 中的验证码
	val, err := redis.Client.Get("code:" + req.Email).Result()
	if err != nil || val != req.Code {
		c.JSON(http.StatusBadRequest, CustomResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid or expired code",
			Data:    nil,
		})
		return
	}

	// 查找用户，不存在则自动注册
	var user model.User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		user = model.User{Email: req.Email}
		db.Create(&user)
	}

	// 生成 JWT token
	accessToken, _ := jwt.GenerateAccessToken(strconv.Itoa(int(user.Uid)), time.Minute*15)
	refreshToken, _ := jwt.GenerateRefreshToken(strconv.Itoa(int(user.Uid)), time.Hour*24*30)
	c.JSON(http.StatusOK, CustomResponse{
		Code:    http.StatusOK,
		Message: "Login successful",
		Data:    gin.H{"accessToken": accessToken, "refreshToken": refreshToken},
	})
}
