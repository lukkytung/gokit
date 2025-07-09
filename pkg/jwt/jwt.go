package jwt

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 获取 JWT 密钥
func getSecretKey() ([]byte, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return nil, jwt.ErrSignatureInvalid
	}
	return []byte(secretKey), nil
}

// 生成access token
func GenerateAccessToken(uid string, duration time.Duration) (string, error) {
	if duration == 0 {
		duration = 15 * time.Minute
	}
	secretKey, err := getSecretKey()
	if err != nil {
		log.Println("Failed to get secret key:", err)
		return "", err
	}
	claims := jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("Failed to create token:", err)
		return "", err
	}
	return tokenStr, nil
}

// 生成refresh token
func GenerateRefreshToken(uid string, duration time.Duration) (string, error) {
	if duration == 0 {
		duration = 7 * 24 * time.Hour
	}
	secretKey, err := getSecretKey()
	if err != nil {
		log.Println("Failed to get secret key:", err)
		return "", err
	}
	claims := jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("Failed to create token:", err)
		return "", err
	}
	return tokenStr, nil
}

// ParseToken 解析 JWT 令牌并返回用户信息
func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		log.Println("Failed to get secret key:", err)
		return nil, err
	}
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
