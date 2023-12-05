package shortener_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bobopylabepolhk/ypshortener/config"
	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener"
	"github.com/stretchr/testify/assert"
)

func TestHandleShortenURL(t *testing.T) {
	t.Run("should save ogURL and respond with shortURL", func(t *testing.T) {
		router := shortener.Router{Us: shortener.NewURLShortenerService()}

		ogURL := "https://lavka.yandex.ru/"
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(ogURL))
		router.HandleShortenURL(rec, req)

		resp := rec.Result()
		assert.Equal(t, resp.StatusCode, http.StatusCreated, "success code should be 201")
		assert.Equal(
			t,
			resp.Header.Get("Content-Type"),
			"text/plain", "Content type header should be text/plain",
		)

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		assert.Nil(t, err)
		assert.Contains(t, string(body), config.APIURL, "should return <APIURL>/<token>")
	})

	t.Run("shoud send 400 code when body isn't provided", func(t *testing.T) {
		router := shortener.Router{Us: shortener.NewURLShortenerService()}

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		router.HandleShortenURL(rec, req)

		resp := rec.Result()
		assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
	})
	t.Run("should create new shortURL if called with same ogURL", func(t *testing.T) {
		router := shortener.Router{Us: shortener.NewURLShortenerService()}

		ogURL := "https://market.yandex.ru/"
		rec := httptest.NewRecorder()

		req1 := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(ogURL))
		router.HandleShortenURL(rec, req1)
		resp1 := rec.Result()

		req2 := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(ogURL))
		router.HandleShortenURL(rec, req2)
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
		assert.Nil(t, err, err)

		router := shortener.Router{Us: us}

		rec := httptest.NewRecorder()
		path := fmt.Sprintf("/%s", token)
		req := httptest.NewRequest(http.MethodGet, path, nil)
		router.HandleGetURL(rec, req)

		resp := rec.Result()
		assert.Equal(
			t,
			resp.StatusCode,
			http.StatusTemporaryRedirect,
			"status code should be 307",
		)
		assert.Equal(
			t,
			resp.Header.Get("Location"),
			ogURL, fmt.Sprintf("Location should be %s", ogURL),
		)

	})
	t.Run("shoud send 404 code when called without saving ogURL first", func(t *testing.T) {
		router := shortener.Router{Us: shortener.NewURLShortenerService()}

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/yU7n23", nil)
		router.HandleGetURL(rec, req)

		resp := rec.Result()
		assert.Equal(
			t,
			resp.StatusCode,
			http.StatusNotFound,
			"status code should be 400",
		)
	})
}
