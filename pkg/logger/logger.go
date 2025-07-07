package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	Log = logger.Sugar()
}
