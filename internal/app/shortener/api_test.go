package shortener_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bobopylabepolhk/ypshortener/config"
	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHandleShortenURL(t *testing.T) {
	t.Run("should save ogURL and respond with shortURL", func(t *testing.T) {
		router := shortener.Router{URLShortenerService: shortener.NewURLShortenerService()}
		ogURL := "https://lavka.yandex.ru/"

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(ogURL))

		e := echo.New()
		ctx := e.NewContext(req, rec)
		router.HandleShortenURL(ctx)

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
		assert.Nil(t, err)
		assert.Contains(t, string(body), config.BASEURL, "should return <BASEURL>/<token>")
	})

	t.Run("shoud send 400 code when body isn't provided", func(t *testing.T) {
		router := shortener.Router{URLShortenerService: shortener.NewURLShortenerService()}

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", nil)

		e := echo.New()
		ctx := e.NewContext(req, rec)
		echoErr := router.HandleShortenURL(ctx)

		assert.Equal(t, http.StatusBadRequest, echoErr.(*echo.HTTPError).Code)
	})

	t.Run("should create new shortURL if called with same ogURL", func(t *testing.T) {
		router := shortener.Router{URLShortenerService: shortener.NewURLShortenerService()}

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
		assert.Nil(t, err)

		defer resp2.Body.Close()
		body2, err := io.ReadAll(resp1.Body)

		assert.Nil(t, err)
		assert.NotEqual(t, string(body1), string(body2))
	})
}

func TestHandleGetURL(t *testing.T) {
	t.Run("should successfully respond with ogURL", func(t *testing.T) {
		us := shortener.NewURLShortenerService()
		token := "Ghf6i9"
		ogURL := "https://yandex.ru/maps/geo/sankt_peterburg/53000093/?ll=30.092322%2C59.940675&z=9.87"
		err := us.SaveShortURL(ogURL, token)
		assert.Nil(t, err)

		router := shortener.Router{URLShortenerService: us}

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

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
		router := shortener.Router{URLShortenerService: shortener.NewURLShortenerService()}

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		e := echo.New()
		ctx := e.NewContext(req, rec)
		ctx.SetParamNames("token")
		ctx.SetParamValues("yU7n23")
		echoErr := router.HandleGetURL(ctx)

		assert.Equal(
			t,
			http.StatusNotFound,
			echoErr.(*echo.HTTPError).Code,
			"status code should be 404",
		)
	})
}
