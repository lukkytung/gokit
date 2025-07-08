package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lukkytung/gokit/pkg/jwt"
)

// AuthMiddleware JWT 鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的 token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			c.Abort()
			return
		}

		// 验证 token
		claims, err := jwt.ParseToken(strings.TrimPrefix(tokenString, "Bearer "))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 将用户信息传递到上下文
		c.Set("uid", claims["uid"])
		log.Fatalf("Authenticated user: %s", claims["uid"])

		c.Next()
	}
}
