package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.SugaredLogger) echo.MiddlewareFunc {
	loggerConfig := middleware.RequestLoggerConfig{
		LogMethod:       true,
		LogResponseSize: true,
		LogLatency:      true,
		LogURI:          true,
		LogStatus:       true,
		LogValuesFunc: func(ctx echo.Context, v middleware.RequestLoggerValues) error {
			logger.Infoln(v.Method, v.URI, v.Status, "response time:", v.Latency, "response size:", v.ResponseSize)
			return nil
		},
	}

	return middleware.RequestLoggerWithConfig(loggerConfig)
}
