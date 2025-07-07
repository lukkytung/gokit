package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/lukkytung/gokit/pkg/config"
	"github.com/lukkytung/gokit/pkg/logger"
)

// Global DB instance
var DB *gorm.DB

// InitPostgres 初始化 PostgreSQL 数据库连接
func InitPostgres() error {
	dbConfig := config.AppConfig

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.DatabaseHost, dbConfig.DatabaseUser, dbConfig.DatabasePassword, dbConfig.DatabaseName, dbConfig.DatabasePort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Errorf("Failed to connect to database: %v", err)
		return err
	}

	DB = db
	logger.Log.Info("PostgreSQL connected successfully")
	return nil
}
