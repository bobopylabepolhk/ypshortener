package logger

import (
	"go.uber.org/zap"

	"github.com/bobopylabepolhk/ypshortener/config"
)

func New() *zap.SugaredLogger {
	logger := zap.Must(zap.NewProduction())

	if config.Cfg.Debug {
		logger = zap.Must(zap.NewDevelopment())
	}

	return logger.Sugar()
}
