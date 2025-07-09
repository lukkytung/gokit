// email.go 负责发送邮件的工具函数
package utils

import (
	"fmt"
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
	pass := dbConfig.EmailPass
	from := dbConfig.EmailFrom
	portStr := dbConfig.EmailSmtpPort

	if host == "" {
		return fmt.Errorf("环境变量 EMAIL_SMTP_HOST 未设置")
	}
	if user == "" {
		return fmt.Errorf("环境变量 EMAIL_USER 未设置")
	}
	if pass == "" {
		return fmt.Errorf("环境变量 EMAIL_PASS 未设置")
	}
	if from == "" {
		return fmt.Errorf("环境变量 EMAIL_FROM 未设置")
	}

	// 解析 SMTP 端口，默认 465
	port := 465 // 默认端口
	if portStr != "" {
		p, err := strconv.Atoi(portStr)
		if err != nil {
			return fmt.Errorf("环境变量 EMAIL_SMTP_PORT 解析失败: %v", err)
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
	pass := dbConfig.EmailPass
	from := dbConfig.EmailFrom
	portStr := dbConfig.EmailSmtpPort

	if host == "" {
		return fmt.Errorf("环境变量 EMAIL_SMTP_HOST 未设置")
	}
	if user == "" {
		return fmt.Errorf("环境变量 EMAIL_USER 未设置")
	}
	if pass == "" {
		return fmt.Errorf("环境变量 EMAIL_PASS 未设置")
	}
	if from == "" {
		return fmt.Errorf("环境变量 EMAIL_FROM 未设置")
	}

	// 解析 SMTP 端口，默认 465
	port := 465 // 默认端口
	if portStr != "" {
		p, err := strconv.Atoi(portStr)
		if err != nil {
			return fmt.Errorf("环境变量 EMAIL_SMTP_PORT 解析失败: %v", err)
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
