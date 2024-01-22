package logger

import (
	"go.uber.org/zap"

	"github.com/bobopylabepolhk/ypshortener/config"
)

var zapLogger *zap.SugaredLogger

func Info(msg any) {
	zapLogger.Info(msg)
}

func New() *zap.SugaredLogger {
	logger := zap.Must(zap.NewProduction())

	if config.Cfg.Debug {
		logger = zap.Must(zap.NewDevelopment())
	}

	zapLogger = logger.Sugar()

	return logger.Sugar()
}
