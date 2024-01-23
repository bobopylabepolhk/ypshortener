package shortener_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener"
)

func TestGetOgURL(t *testing.T) {
	t.Run("should return err if ogURL was never saved", func(t *testing.T) {
		repo := shortener.NewURLShortenerRepo()
		us := shortener.NewURLShortenerService(repo)

		_, err := us.GetOriginalURL("blahblah")
		assert.Error(t, err)
	})

	t.Run("should return ogURL, nil for a given token", func(t *testing.T) {
		repo := shortener.NewURLShortenerRepo()
		us := shortener.NewURLShortenerService(repo)

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
	repo := shortener.NewURLShortenerRepo()
	us := shortener.NewURLShortenerService(repo)

	t.Run("should return err if called with invalid url", func(t *testing.T) {
		assert.Error(t, us.SaveShortURL("blahblah", "YUG76a"), t.Name())
	})
}
