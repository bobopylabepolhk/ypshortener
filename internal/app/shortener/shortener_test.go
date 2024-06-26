package shortener_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener"
	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener/repo"
)

func TestGetOgURL(t *testing.T) {
	t.Run("should return err if ogURL was never saved", func(t *testing.T) {
		repo, err := repo.NewURLShortenerRepo()
		assert.NoError(t, err)
		us := shortener.NewURLShortenerService(repo)

		_, err = us.GetOriginalURL(context.Background(), "blahblah")
		assert.Error(t, err)
	})

	t.Run("should return ogURL, nil for a given token", func(t *testing.T) {
		repo, err := repo.NewURLShortenerRepo()
		assert.NoError(t, err)
		us := shortener.NewURLShortenerService(repo)

		token := "6Tg8oJ"
		ogURL := "https://yandex.com/"
		_, err = us.SaveShortURL(context.Background(), ogURL, token)
		require.NoError(t, err)
		r, err := us.GetOriginalURL(context.Background(), token)
		assert.NoError(t, err)
		assert.Equal(t, ogURL, r)
	})

}

func TestSaveShortURL(t *testing.T) {
	repo, err := repo.NewURLShortenerRepo()
	assert.NoError(t, err)
	us := shortener.NewURLShortenerService(repo)

	t.Run("should return err if called with invalid url", func(t *testing.T) {
		_, err := us.SaveShortURL(context.Background(), "blahblah", "YUG76a")
		assert.Error(t, err, t.Name())
	})
}
