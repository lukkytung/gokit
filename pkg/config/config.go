package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// 配置结构体
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Db       int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// 全局配置变量
var AppConfig Config

// InitConfig 加载配置文件
func InitConfig(configPath string) error {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return err
	}

	// 反序列化配置文件内容到 Config 结构体
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return err
	}

	// 打印加载的配置，便于调试
	fmt.Printf("Loaded config: %+v\n", AppConfig)
	return nil
}
