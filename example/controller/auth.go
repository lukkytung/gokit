package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lukkytung/gokit/example/model"
	"github.com/lukkytung/gokit/pkg/jwt"
	"github.com/lukkytung/gokit/pkg/service"
	"github.com/lukkytung/gokit/pkg/utils"
)

type CustomResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type SendEmailRequest struct {
	Email string `json:"email"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type UserResponse struct {
	Uid       string    `json:"uid"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func SendCode(c *gin.Context) {
	var req SendEmailRequest
	// 解析请求体中的 email
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, CustomResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid email",
			Data:    err.Error(),
		})
		return
	}

	// 生成 6 位随机验证码
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	// 将验证码存入 Redis，有效期 10 分钟
	err := service.RedisClient.Set("code:"+req.Email, code, time.Minute*10).Err()
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

	db := service.DB

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
	val, err := service.RedisClient.Get("code:" + req.Email).Result()
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
		// 用户不存在,创建新用户
		user = model.User{Email: req.Email}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, CustomResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to create user",
				Data:    err.Error(),
			})
			return
		}
	}

	// 生成 JWT token
	at, rt, _, err := jwt.GenerateTokens(user.Uid, time.Minute*15, time.Hour*24*30)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CustomResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate tokens",
			Data:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, CustomResponse{
		Code:    http.StatusOK,
		Message: "Login successful",
		Data:    gin.H{"accessToken": at, "refreshToken": rt, "user": UserResponse{Uid: user.Uid, Email: user.Email, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}},
	})
}

func RefreshToken(c *gin.Context) {
	refreshToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, CustomResponse{
			Code:    http.StatusBadRequest,
			Message: "Missing refresh token",
			Data:    nil,
		})
		return
	}

	claims, err := jwt.ParseToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, CustomResponse{
			Code:    http.StatusUnauthorized,
			Message: "Invalid refresh token",
			Data:    nil,
		})
		return
	}

	jtiKey := "refresh_jti:" + claims.JTI
	// 检查 JTI 是否存在
	_, err = service.RedisClient.Get(jtiKey).Result()
	if err != nil {
		c.JSON(http.StatusUnauthorized, CustomResponse{
			Code:    http.StatusUnauthorized,
			Message: "Refresh token expired",
			Data:    nil,
		})
		return
	}
	// 生成新的 token
	at, rt, _, err := jwt.GenerateTokens(claims.Uid, time.Minute*15, time.Hour*24*30)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CustomResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate tokens",
			Data:    err.Error(),
		})
		return
	}

	// 返回新的 token
	c.JSON(http.StatusOK, CustomResponse{
		Code:    http.StatusOK,
		Message: "Refresh token successful",
		Data:    gin.H{"access_token": at, "refresh_token": rt},
	})
}

func LogoutHandler(c *gin.Context) {
	jti := c.GetString("jti")

	// 删除 Redis 中对应的 refresh_jti
	service.RedisClient.Del("refresh_jti:" + jti)

	c.JSON(http.StatusOK, CustomResponse{
		Code:    http.StatusOK,
		Message: "Logged out successfully",
		Data:    nil,
	})
}
