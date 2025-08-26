// email.go 负责发送邮件的工具函数
package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/lukkytung/gokit/pkg/config"
	"gopkg.in/gomail.v2"
)

// SendEmailWithTemplate 使用环境变量中的 SMTP 配置发送邮件
// to: 收件人邮箱，subject: 邮件主题，body: 邮件内容
func SendEmailWithTemplate(to string, subject string, body string) error {
	dbConfig := config.AppConfig

	// 检查环境变量
	host := dbConfig.EmailSmtpHost
	user := dbConfig.EmailUser
	pass := dbConfig.EmailPassword
	from := dbConfig.EmailFrom
	portStr := dbConfig.EmailSmtpPort

	if host == "" {
		log.Printf("环境变量 EMAIL_SMTP_HOST 未设置")
		return nil
	}
	if user == "" {
		log.Printf("环境变量 EMAIL_USER 未设置")
		return nil
	}
	if pass == "" {
		log.Printf("环境变量 EMAIL_PASSWORD 未设置")
		return nil
	}
	if from == "" {
		log.Printf("环境变量 EMAIL_FROM 未设置")
		return nil
	}

	// 解析 SMTP 端口，默认 465
	port := 465 // 默认端口
	if portStr != "" {
		p, err := strconv.Atoi(portStr)
		if err != nil {
			log.Printf("环境变量 EMAIL_SMTP_PORT 解析失败: %v", err)
			return err
		}
		port = p
	}

	// 构建邮件消息
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// 创建 SMTP 拨号器
	d := gomail.NewDialer(
		host,
		port,
		user,
		pass,
	)

	// 发送邮件
	return d.DialAndSend(m)
}

// SendEmail 使用环境变量中的 SMTP 配置发送邮件
// to: 收件人邮箱，subject: 邮件主题，body: 邮件内容
func SendEmail(to string, subject string, body string) error {
	dbConfig := config.AppConfig

	// 检查环境变量
	host := dbConfig.EmailSmtpHost
	user := dbConfig.EmailUser
	pass := dbConfig.EmailPassword
	from := dbConfig.EmailFrom
	portStr := dbConfig.EmailSmtpPort

	if host == "" {
		log.Printf("环境变量 EMAIL_SMTP_HOST 未设置")
		return nil
	}
	if user == "" {
		log.Printf("环境变量 EMAIL_USER 未设置")
	}
	if pass == "" {
		log.Printf("环境变量 EMAIL_PASSWORD 未设置")
		return nil
	}
	if from == "" {
		log.Printf("环境变量 EMAIL_FROM 未设置")
		return nil
	}

	// 解析 SMTP 端口，默认 465
	port := 465 // 默认端口
	if portStr != "" {
		p, err := strconv.Atoi(portStr)
		if err != nil {
			log.Printf("环境变量 EMAIL_SMTP_PORT 解析失败: %v", err)
			return err
		}
		port = p
	}

	// 构建邮件消息
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// 创建 SMTP 拨号器
	d := gomail.NewDialer(
		host,
		port,
		user,
		pass,
	)

	// 发送邮件
	return d.DialAndSend(m)
}

// 获取 AES Key
func getAESKey() ([]byte, error) {
	keyHex := config.AppConfig.EmailEncrypSecretKey
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid secret key hex: %w", err)
	}
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, fmt.Errorf("invalid AES key length: %d", len(key))
	}
	return key, nil
}

// EncryptEmailDeterministic 确定性加密
func EncryptEmailDeterministic(email string) (string, error) {
	key, err := getAESKey()
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize()) // 固定 nonce
	cipherText := aesGCM.Seal(nil, nonce, []byte(email), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// DecryptEmailDeterministic 解密
func DecryptEmailDeterministic(encrypted string) (string, error) {
	key, err := getAESKey()
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize()) // 固定 nonce
	plainText, err := aesGCM.Open(nil, nonce, data, nil)
	if err != nil {
		return "", errors.New("failed to decrypt email")
	}

	return string(plainText), nil
}
