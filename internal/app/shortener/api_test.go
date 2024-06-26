package shortener_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bobopylabepolhk/ypshortener/config"
	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener"
	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener/repo"
)

func TestHandleShortenURL(t *testing.T) {
	t.Run("should save ogURL and respond with shortURL", func(t *testing.T) {
		repo, err := repo.NewURLShortenerRepo()
		assert.NoError(t, err)
		router := shortener.Router{URLShortenerService: shortener.NewURLShortenerService(repo)}
		ogURL := "https://lavka.yandex.ru/"

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(ogURL))

		e := echo.New()
		ctx := e.NewContext(req, rec)
		err = router.HandleShortenURL(ctx)
		require.NoError(t, err)

		resp := rec.Result()
		assert.Equal(t, http.StatusCreated, resp.StatusCode, "success code should be 201")
		assert.Contains(
			t,
			resp.Header.Get("Content-Type"),
			"text/plain",
			"Content type header should be text/plain",
		)

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Contains(t, string(body), config.Cfg.BaseURL, "should return <BASEURL>/<token>")
	})

	t.Run("shoud send 400 code when body isn't provided", func(t *testing.T) {
		repo, err := repo.NewURLShortenerRepo()
		assert.NoError(t, err)
		router := shortener.Router{URLShortenerService: shortener.NewURLShortenerService(repo)}

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", nil)

		e := echo.New()
		ctx := e.NewContext(req, rec)
		err = router.HandleShortenURL(ctx)

		assert.EqualError(t, err, echo.ErrBadRequest.Error())
	})

	t.Run("should create new shortURL if called with same ogURL", func(t *testing.T) {
		repo, err := repo.NewURLShortenerRepo()
		assert.NoError(t, err)
		router := shortener.Router{URLShortenerService: shortener.NewURLShortenerService(repo)}

		ogURL := "https://market.yandex.ru/"
		rec := httptest.NewRecorder()
		e := echo.New()

		req1 := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(ogURL))
		router.HandleShortenURL(e.NewContext(req1, rec))
		resp1 := rec.Result()

		req2 := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(ogURL))
		router.HandleShortenURL(e.NewContext(req2, rec))
		resp2 := rec.Result()

		defer resp1.Body.Close()
		body1, err := io.ReadAll(resp1.Body)
		assert.NoError(t, err)

		defer resp2.Body.Close()
		body2, err := io.ReadAll(resp2.Body)

		assert.NoError(t, err)
		assert.NotEqual(t, string(body1), string(body2))
	})
}

func TestHandleJSONShortenURl(t *testing.T) {
	t.Run("shoud send 422 code when body isn't provided", func(t *testing.T) {
		repo, err := repo.NewURLShortenerRepo()
		assert.NoError(t, err)
		router := shortener.Router{URLShortenerService: shortener.NewURLShortenerService(repo)}

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/api/shorten", nil)

		e := echo.New()
		ctx := e.NewContext(req, rec)
		err = router.HandleJSONShortenURL(ctx)

		assert.EqualError(t, err, echo.ErrUnprocessableEntity.Error())
	})

	t.Run("shoud send 422 code when json is missing url field", func(t *testing.T) {
		repo, err := repo.NewURLShortenerRepo()
		assert.NoError(t, err)
		router := shortener.Router{URLShortenerService: shortener.NewURLShortenerService(repo)}
		e := echo.New()

		tests := []map[string]interface{}{
			{
				"urll": "https://practicum.yandex.ru",
			},
			{
				"url": "",
			},
			{},
		}
		for _, test := range tests {
			data, _ := json.Marshal(test)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewReader(data))

			ctx := e.NewContext(req, rec)
			err := router.HandleJSONShortenURL(ctx)

			assert.EqualError(t, err, echo.ErrUnprocessableEntity.Error())
		}
	})
}

func TestHandleGetURL(t *testing.T) {
	t.Run("should successfully respond with ogURL", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		repo, err := repo.NewURLShortenerRepo()
		assert.NoError(t, err)
		us := shortener.NewURLShortenerService(repo)
		token := "Ghf6i9"
		ogURL := "https://yandex.ru/maps/geo/sankt_peterburg/53000093/?ll=30.092322%2C59.940675&z=9.87"
		_, err = us.SaveShortURL(req.Context(), ogURL, token)
		assert.NoError(t, err)

		router := shortener.Router{URLShortenerService: us}

		e := echo.New()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("token")
		ctx.SetParamValues(token)
		router.HandleGetURL(ctx)

		resp := rec.Result()
		defer resp.Body.Close()
		assert.Equal(
			t,
			http.StatusTemporaryRedirect,
			resp.StatusCode,
			"status code should be 307",
		)
		assert.Equal(
			t,
			ogURL,
			resp.Header.Get("Location"),
			"Location should be set",
		)
	})

	t.Run("shoud send 404 code when called without saving ogURL first", func(t *testing.T) {
		repo, err := repo.NewURLShortenerRepo()
		assert.NoError(t, err)
		router := shortener.Router{URLShortenerService: shortener.NewURLShortenerService(repo)}

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		e := echo.New()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("token")
		ctx.SetParamValues("yU7n23")
		err = router.HandleGetURL(ctx)

		assert.EqualError(t, err, echo.ErrNotFound.Error())
	})
}
