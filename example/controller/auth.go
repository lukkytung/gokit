package controller

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lukkytung/gokit/example/model"
	"github.com/lukkytung/gokit/pkg/database"
	"github.com/lukkytung/gokit/pkg/jwt"
	"github.com/lukkytung/gokit/pkg/redis"
	"github.com/lukkytung/gokit/pkg/utils"
)

// 创建一个全局 context 用于 Redis 操作
var ctx = context.Background()

type SendEmailRequest struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func SendCode(c *gin.Context) {
	type Req struct {
		Email string `json:"email"`
	}
	var req Req
	// 解析请求体中的 email
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid email"})
		return
	}

	// 生成 6 位随机验证码
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	// 将验证码存入 Redis，有效期 10 分钟
	err := redis.Client.Set(ctx, "code:"+req.Email, code, time.Minute*10).Err()
	if err != nil {
		c.JSON(500, gin.H{"error": "Redis error", "detail": err.Error()})
		return
	}

	// 发送邮件（忽略发送失败）
	err = utils.SendEmail(req.Email, "Your Login Code", "Your code is: "+code)
	if err != nil {
		c.JSON(500, gin.H{"error": "Email sending failed"})
		return
	}

	c.JSON(200, gin.H{"message": "Code sent", "code": code})
}

func LoginWithCode(c *gin.Context) {

	db := database.DB

	var req LoginRequest
	// 解析请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// 校验 Redis 中的验证码
	val, err := redis.Client.Get(ctx, "code:"+req.Email).Result()
	if err != nil || val != req.Code {
		c.JSON(400, gin.H{"error": "Invalid or expired code"})
		return
	}

	// 查找用户，不存在则自动注册
	var user model.User
	if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		user = model.User{Email: req.Email}
		db.Create(&user)
		fmt.Println("user created", user)
	}

	// 生成 JWT token
	accessToken, _ := jwt.GenerateAccessToken(strconv.FormatUint(user.Uid, 10), time.Minute*15)
	refreshToken, _ := jwt.GenerateRefreshToken(strconv.FormatUint(user.Uid, 10), time.Hour*24*30)
	c.JSON(200, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}
