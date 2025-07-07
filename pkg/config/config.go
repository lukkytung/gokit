package config

import (
	"github.com/spf13/viper"
)

func InitConfig(configPath string) error {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
