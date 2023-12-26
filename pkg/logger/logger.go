package logger

import (
	"github.com/bobopylabepolhk/ypshortener/config"
	"go.uber.org/zap"
)

func New() *zap.SugaredLogger {
	logger := zap.Must(zap.NewProduction())

	if config.Cfg.Debug {
		logger = zap.Must(zap.NewDevelopment())
	}

	return logger.Sugar()
}
