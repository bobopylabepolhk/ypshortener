package shortener

import (
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bobopylabepolhk/ypshortener/config"
	"github.com/bobopylabepolhk/ypshortener/pkg/urlutils"
)

type (
	URLShortener interface {
		SaveShortURL(url string, token string) error
		GetOriginalURL(shortURL string) (string, error)
	}

	Router struct {
		URLShortenerService URLShortener
	}
)

func (router *Router) HandleGetURL(ctx echo.Context) error {
	token := ctx.Param("token")

	ogURL, err := router.URLShortenerService.GetOriginalURL(token)
	if err != nil {
		return echo.ErrNotFound
	}

	ctx.Response().Header().Add("Location", ogURL)
	return ctx.NoContent(http.StatusTemporaryRedirect)
}

func (router *Router) HandleShortenURL(ctx echo.Context) error {
	ogURL, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return echo.ErrBadRequest
	}

	token := urlutils.GetShortURLToken()
	err = router.URLShortenerService.SaveShortURL(string(ogURL), token)

	if err != nil {
		return echo.ErrBadRequest
	}

	res := fmt.Sprintf("%s/%s", config.Cfg.BaseURL, token)
	return ctx.String(http.StatusCreated, res)
}

func NewRouter(e *echo.Echo) {
	us := NewURLShortenerService()
	router := &Router{URLShortenerService: us}
	e.GET("/:token", router.HandleGetURL)
	e.POST("/", router.HandleShortenURL)
}
