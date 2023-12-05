package shortener

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bobopylabepolhk/ypshortener/config"
	"github.com/labstack/echo/v4"
)

type (
	URLShortener interface {
		GetShortURLToken() string
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

	token := router.URLShortenerService.GetShortURLToken()
	err = router.URLShortenerService.SaveShortURL(string(ogURL), token)

	if err != nil {
		return echo.ErrBadRequest
	}

	res := fmt.Sprintf("%s/%s", config.APIURL, token)
	return ctx.String(http.StatusCreated, res)
}

func NewRouter(e *echo.Echo) {
	us := NewURLShortenerService()
	router := &Router{URLShortenerService: us}
	e.GET("/:token", router.HandleGetURL)
	e.POST("/", router.HandleShortenURL)
}
