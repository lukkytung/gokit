package config

import (
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// 配置结构体
type Config struct {
	ServerPort       string
	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	RedisHost        string
	RedisPort        int
	RedisPassword    string
	RedisDb          int
	JwtSecretKey     string
	EmailSmtpHost    string
	EmailUser        string
	EmailPass        string
	EmailFrom        string
	EmailSmtpPort    string
}

// 全局配置变量
var AppConfig Config

// InitConfig 加载配置文件
func InitConfig() error {
	// 只加载 .env 文件（如果存在）
	_ = godotenv.Load()

	// 支持从环境变量读取配置
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// 手动绑定每个字段
	viper.BindEnv("ServerPort", "SERVER_PORT")
	viper.BindEnv("DatabaseHost", "DATABASE_HOST")
	viper.BindEnv("DatabasePort", "DATABASE_PORT")
	viper.BindEnv("DatabaseUser", "DATABASE_USER")
	viper.BindEnv("DatabasePassword", "DATABASE_PASSWORD")
	viper.BindEnv("DatabaseName", "DATABASE_NAME")
	viper.BindEnv("RedisHost", "REDIS_HOST")
	viper.BindEnv("RedisPort", "REDIS_PORT")
	viper.BindEnv("RedisPassword", "REDIS_PASSWORD")
	viper.BindEnv("RedisDb", "REDIS_DB")
	viper.BindEnv("JwtSecretKey", "JWT_SECRET_KEY")
	viper.BindEnv("EmailSmtpHost", "EMAIL_SMTP_HOST")
	viper.BindEnv("EmailUser", "EMAIL_USER")
	viper.BindEnv("EmailPass", "EMAIL_PASS")
	viper.BindEnv("EmailFrom", "EMAIL_FROM")
	viper.BindEnv("EmailSmtpPort", "EMAIL_SMTP_PORT")

	// 反序列化环境变量到 Config 结构体
	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return err
	}

	// 打印加载的配置，便于调试
	log.Printf("Loaded config: %+v\n", AppConfig)

	return nil
}
