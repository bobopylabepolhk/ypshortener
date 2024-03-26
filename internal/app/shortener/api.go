package shortener

import (
	"database/sql"
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener/repo"
	"github.com/bobopylabepolhk/ypshortener/pkg/logger"
	"github.com/bobopylabepolhk/ypshortener/pkg/urlutils"
)

type (
	URLShortener interface {
		SaveShortURL(url string, token string) (string, error)
		GetOriginalURL(shortURL string) (string, error)
		SaveURLBatch(batch []ShortenBatchRequestDTO) ([]ShortenBatchResponseDTO, error)
		GetExistingShortURL(ogURL string) (string, error)
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

	ogURLStr := string(ogURL)
	token := urlutils.GetShortURLToken()
	res, err := router.URLShortenerService.SaveShortURL(ogURLStr, token)
	if err != nil {
		if errors.Is(err, repo.ErrDuplicateURL) {
			shortURL, err := router.URLShortenerService.GetExistingShortURL(ogURLStr)
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
	token := urlutils.GetShortURLToken()
	shortURL, err := router.URLShortenerService.SaveShortURL(ogURLStr, token)

	if err != nil {
		if errors.Is(err, repo.ErrDuplicateURL) {
			shortURL, err := router.URLShortenerService.GetExistingShortURL(ogURLStr)
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

	res, err := router.URLShortenerService.SaveURLBatch(data)

	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusCreated, res)
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
}
