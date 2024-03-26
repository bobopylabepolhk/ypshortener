package logger

import (
	"go.uber.org/zap"

	"github.com/bobopylabepolhk/ypshortener/config"
)

var zapLogger *zap.Logger

func New() *zap.SugaredLogger {
	logger := zap.Must(zap.NewProduction())

	if config.Cfg.Debug {
		logger = zap.Must(zap.NewDevelopment())
	}

	zapLogger = logger

	return logger.Sugar()
}

func Debug(m string, v ...zap.Field) {
	zapLogger.Debug(m, v...)
}

func Info(m string, v ...zap.Field) {
	zapLogger.Debug(m, v...)
}

func Error(m string, v ...zap.Field) {
	zapLogger.Debug(m, v...)
}

func Fatal(m string, v ...zap.Field) {
	zapLogger.Debug(m, v...)
}
