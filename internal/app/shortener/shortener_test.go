package shortener_test

import (
	"regexp"
	"testing"

	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetToken(t *testing.T) {
	us := shortener.NewURLShortenerService()
	iterations := 99

	t.Run("token should never contain: ? / # & % . ,", func(t *testing.T) {
		for i := 0; i < iterations; i++ {
			token := us.GetShortURLToken()
			r := regexp.MustCompile(`[^?/#&%.,]*$`)

			require.Regexp(t, r, token)
		}
	})
}

func TestGetOgURL(t *testing.T) {
	t.Run("should return err if ogURL was never saved", func(t *testing.T) {
		us := shortener.NewURLShortenerService()

		_, err := us.GetOriginalURL("blahblah")
		assert.Error(t, err)
	})

	t.Run("should return ogURL, nil for a given token", func(t *testing.T) {
		us := shortener.NewURLShortenerService()

		token := "6Tg8oJ"
		ogURL := "https://yandex.com/"
		err := us.SaveShortURL(ogURL, token)
		require.NoError(t, err)
		r, err := us.GetOriginalURL(token)
		assert.NoError(t, err)
		assert.Equal(t, ogURL, r)
	})

}

func TestSaveShortURL(t *testing.T) {
	us := shortener.NewURLShortenerService()

	t.Run("should return err if called with invalid url", func(t *testing.T) {
		assert.Error(t, us.SaveShortURL("blahblah", "YUG76a"), t.Name())
	})
}
