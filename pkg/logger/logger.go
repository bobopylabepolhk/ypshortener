package logger

import (
	"go.uber.org/zap"

	"github.com/bobopylabepolhk/ypshortener/config"
)

var zapLogger *zap.SugaredLogger

func New() *zap.SugaredLogger {
	logger := zap.Must(zap.NewProduction())

	if config.Cfg.Debug {
		logger = zap.Must(zap.NewDevelopment())
	}

	zapLogger = logger.Sugar()

	return zapLogger
}

func Debug(args ...interface{}) {
	zapLogger.Debug(args)
}

func Info(args ...interface{}) {
	zapLogger.Debug(args)
}

func Error(args ...interface{}) {
	zapLogger.Debug(args)
}

func Fatal(args ...interface{}) {
	zapLogger.Debug(args)
}
