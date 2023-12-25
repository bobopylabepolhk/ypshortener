package logger

import (
	"github.com/bobopylabepolhk/ypshortener/config"
	"go.uber.org/zap"
)

func New() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()

	if config.Cfg.Debug {
		logger, err = zap.NewDevelopment()
	}

	return logger.Sugar(), err
}
