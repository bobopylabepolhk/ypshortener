package shortener

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener/repo"
	"github.com/bobopylabepolhk/ypshortener/pkg/auth"
	"github.com/bobopylabepolhk/ypshortener/pkg/logger"
	"github.com/bobopylabepolhk/ypshortener/pkg/urlutils"
)

type (
	URLShortener interface {
		SaveShortURL(ctx context.Context, url string, token string, userID string) (string, error)
		GetOriginalURL(ctx context.Context, shortURL string) (string, error)
		SaveURLBatch(ctx context.Context, batch []ShortenBatchRequestDTO, userID string) ([]ShortenBatchResponseDTO, error)
		GetExistingShortURL(ctx context.Context, ogURL string) (string, error)
		GetUserURLs(ctx context.Context, userID string) ([]repo.URLBatch, error)
	}

	Router struct {
		URLShortenerService URLShortener
	}
)

func (router *Router) HandleGetURL(ctx echo.Context) error {
	token := ctx.Param("token")

	ogURL, err := router.URLShortenerService.GetOriginalURL(ctx.Request().Context(), token)
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

	ogURLStr := string(ogURL)
	token := urlutils.CreateRandomToken(6)
	res, err := router.URLShortenerService.SaveShortURL(ctx.Request().Context(), ogURLStr, token, auth.GetUserID(ctx))
	if err != nil {
		if errors.Is(err, repo.ErrDuplicateURL) {
			shortURL, err := router.URLShortenerService.GetExistingShortURL(ctx.Request().Context(), ogURLStr)
			if err != nil {
				return echo.ErrInternalServerError
			}

			return ctx.String(http.StatusConflict, shortURL)
		}

		return echo.ErrBadRequest
	}

	return ctx.String(http.StatusCreated, res)
}

type ShortenURLRequestDTO struct {
	URL string `json:"url"`
}

type ShortenURLResponseDTO struct {
	Result string `json:"result"`
}

func (router *Router) HandleJSONShortenURL(ctx echo.Context) error {
	data := new(ShortenURLRequestDTO)
	err := ctx.Bind(&data)

	if err != nil || data.URL == "" {
		return echo.ErrUnprocessableEntity
	}

	ogURLStr := data.URL
	token := urlutils.CreateRandomToken(6)
	shortURL, err := router.URLShortenerService.SaveShortURL(ctx.Request().Context(), ogURLStr, token, auth.GetUserID(ctx))

	if err != nil {
		if errors.Is(err, repo.ErrDuplicateURL) {
			shortURL, err := router.URLShortenerService.GetExistingShortURL(ctx.Request().Context(), ogURLStr)
			if err != nil {
				return echo.ErrInternalServerError
			}

			return ctx.JSON(http.StatusConflict, ShortenURLResponseDTO{Result: shortURL})
		}

		return echo.ErrBadRequest
	}

	res := ShortenURLResponseDTO{Result: shortURL}

	return ctx.JSON(http.StatusCreated, res)
}

type ShortenBatchRequestDTO struct {
	CorrelationID string `json:"correlation_id"`
	OgURL         string `json:"original_url"`
}

type ShortenBatchResponseDTO struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

func (router *Router) HandleBatchShortenURL(ctx echo.Context) error {
	data := make([]ShortenBatchRequestDTO, 0)
	err := ctx.Bind(&data)

	if err != nil {
		return echo.ErrUnprocessableEntity
	}

	res, err := router.URLShortenerService.SaveURLBatch(ctx.Request().Context(), data, auth.GetUserID(ctx))

	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (router *Router) HandleUserUrls(ctx echo.Context) error {
	cookie, err := ctx.Cookie(auth.USER_ID_COOKIE)
	if err != nil {
		return echo.ErrUnauthorized
	}

	res, err := router.URLShortenerService.GetUserURLs(ctx.Request().Context(), cookie.Value)
	if err != nil {
		return echo.ErrUnauthorized
	}

	if len(res) > 0 {
		return ctx.JSON(http.StatusOK, res)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func NewRouter(e *echo.Echo, db *sql.DB) {
	repo, err := repo.NewURLShortenerRepo(repo.WithPostgres(db))
	if err != nil {
		logger.Fatal("failed to initialize URLShortenerRepository")
	}
	us := NewURLShortenerService(repo)
	router := &Router{URLShortenerService: us}
	e.GET("/:token", router.HandleGetURL)
	e.POST("/", router.HandleShortenURL)

	e.POST("/api/shorten", router.HandleJSONShortenURL)
	e.POST("/api/shorten/batch", router.HandleBatchShortenURL)
	e.GET("/api/user/urls", router.HandleUserUrls)
}
