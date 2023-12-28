package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GzipMiddleware() echo.MiddlewareFunc {
	return middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(ctx echo.Context) bool {
			contentType := ctx.Request().Header.Get("Content-Type")
			compressableTypes := []string{"application/json", "text/html"}

			for _, t := range compressableTypes {
				if contentType == t {
					return false
				}
			}

			return true
		},
	})
}
