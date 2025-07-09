package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lukkytung/gokit/pkg/jwt"
	"github.com/lukkytung/gokit/pkg/service"
)

// AuthMiddleware JWT 鉴权中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		claims, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		// 校验是否被拉黑
		jtiKey := "refresh_jti:" + claims.JTI
		_, err = service.RedisClient.Get(jtiKey).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token revoked"})
			return
		}

		c.Set("uid", claims.Uid)
		c.Set("jti", claims.JTI)
		c.Next()
	}
}
