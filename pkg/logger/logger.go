package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 全局日志变量
var Log *zap.SugaredLogger

// InitLogger 初始化 Zap Logger，设置为开发模式
func InitLogger() {
	// 设置日志级别和输出格式
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	Log = logger.Sugar()
	Log.Info("Logger initialized")
}

// Shutdown 用于优雅关闭日志
func Shutdown() {
	if err := Log.Sync(); err != nil && !strings.Contains(err.Error(), "inappropriate ioctl for device") {
		Log.Errorf("Failed to sync logger: %v", err)
	}
}
