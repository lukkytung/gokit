package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken 生成 JWT 令牌
func GenerateToken(uid string, duration time.Duration) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	claims := jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ParseToken 解析 JWT 令牌并返回用户信息
func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
