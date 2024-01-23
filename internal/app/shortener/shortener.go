package shortener

import (
	"errors"

	"github.com/bobopylabepolhk/ypshortener/pkg/urlutils"
)

type (
	URLShortenerRepository interface {
		CreateShortURL(token string, ogURL string)
		GetOgURL(shortURL string) (string, error)
	}

	URLShortenerService struct {
		repo URLShortenerRepository
	}
)

func NewURLShortenerService(repo *URLShortenerRepo) *URLShortenerService {
	return &URLShortenerService{
		repo: repo,
	}
}

func (us URLShortenerService) SaveShortURL(url string, token string) error {
	if !urlutils.ValidateURL(url) {
		return errors.New("not a valid url")
	}

	us.repo.CreateShortURL(token, url)

	return nil
}

func (us URLShortenerService) GetOriginalURL(shortURL string) (string, error) {
	return us.repo.GetOgURL(shortURL)
}
