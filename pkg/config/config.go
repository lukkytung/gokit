package config

import (
	"log"
	"os"
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
	EmailPassword    string
	EmailFrom        string
	EmailSmtpPort    string
}

// 全局配置变量
var AppConfig Config

// InitConfig 加载配置文件
func InitConfig() error {
	// 获取当前的环境（默认开发环境）
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development" // 默认为开发环境
	}

	// 根据环境加载相应的 .env 文件
	var envFile string
	switch env {
	case "production":
		envFile = ".env.production"
	case "staging":
		envFile = ".env.staging"
	default:
		envFile = ".env" // 默认加载开发环境的配置
	}

	// 加载 .env 文件
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file, %s", err)
		return err
	}
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
	viper.BindEnv("EmailPassword", "EMAIL_PASSWORD")
	viper.BindEnv("EmailFrom", "EMAIL_FROM")
	viper.BindEnv("EmailSmtpPort", "EMAIL_SMTP_PORT")

	// 反序列化环境变量到 Config 结构体
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return err
	}

	// 打印加载的配置，便于调试
	log.Printf("Loaded config: %+v\n", AppConfig)

	return nil
}
