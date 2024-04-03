package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bobopylabepolhk/ypshortener/pkg/auth"
)

func AuthMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			c, err := ctx.Cookie(auth.UserIDCookie)
			if err != nil || !auth.ValidateUserID(c.Value, secret) {
				c = new(http.Cookie)
				c.Name = auth.UserIDCookie
				userID := auth.GenerateUserID(secret)
				c.Value = userID
				ctx.SetCookie(c)
				ctx.Set(auth.UserIDCookie, userID)
			}

			return next(ctx)
		}
	}
}
