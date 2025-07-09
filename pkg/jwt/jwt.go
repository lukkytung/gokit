package jwt

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lukkytung/gokit/pkg/redis"
	"github.com/lukkytung/gokit/pkg/utils"
)

type Claims struct {
	Uid string `json:"uid"`
	JTI string `json:"jti"`
	jwt.RegisteredClaims
}

// 获取 JWT 密钥
func getSecretKey() ([]byte, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return nil, jwt.ErrSignatureInvalid
	}
	return []byte(secretKey), nil
}

func GenerateTokens(uid string, accessDuration time.Duration, refreshDuration time.Duration) (accessToken, refreshToken, jti string, err error) {
	jti, err = utils.GenerateID()
	if err != nil {
		return
	}

	if accessDuration == 0 {
		accessDuration = 15 * time.Minute
	}
	if refreshDuration == 0 {
		refreshDuration = 7 * 24 * time.Hour
	}

	accessClaims := Claims{
		Uid: uid,
		JTI: jti,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessDuration)),
		},
	}

	secretKey, err := getSecretKey()
	if err != nil {
		log.Println("Failed to get secret key:", err)
		return "", "", "", err
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = at.SignedString(secretKey)
	if err != nil {
		return "", "", "", err
	}

	refreshClaims := Claims{
		Uid: uid,
		JTI: jti,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshDuration)),
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = rt.SignedString(secretKey)
	if err != nil {
		return "", "", "", err
	}

	// 存 JTI 到 Redis，用于刷新 token 时校验，超时后自动删除，需要重新登录
	redis.Client.Set("refresh_jti:"+jti, uid, refreshDuration)
	return accessToken, refreshToken, jti, err
}

// ParseToken 解析 JWT 令牌并返回用户信息
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getSecretKey()
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*Claims), nil
}
