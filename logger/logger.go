package logger

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	// sugar := logger.Sugar()
}

// Info
//
//	@Description: info日志
//	@param msg 信息
//	@param fields 日志字段
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Error
//
//	@Description: Error日志
//	@param msg 信息
//	@param fields 日志字段
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}
