// email.go 负责发送邮件的工具函数
package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
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

var secretKey = []byte("12345678901234567890123456789012") // 32字节 = AES-256

// 加密邮箱
func EncryptEmail(email string) (string, error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(email), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// 解密邮箱
func DecryptEmail(encrypted string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
