package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 生成access token
func GenerateAccessToken(uid string, duration time.Duration) (string, error) {
	if duration == 0 {
		duration = 15 * time.Minute
	}
	secretKey := os.Getenv("JWT_SECRET_KEY")
	claims := jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// 生成refresh token
func GenerateRefreshToken(uid string, duration time.Duration) (string, error) {
	if duration == 0 {
		duration = 7 * 24 * time.Hour
	}
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
