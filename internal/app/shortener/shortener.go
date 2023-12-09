package shortener

import (
	"errors"
	"fmt"

	"github.com/bobopylabepolhk/ypshortener/pkg/urlutils"
)

type URLShortenerService struct {
	urls map[string]string
}

func NewURLShortenerService() *URLShortenerService {
	return &URLShortenerService{
		urls: make(map[string]string),
	}
}

func (us URLShortenerService) SaveShortURL(url string, token string) error {
	if !urlutils.ValidateURL(url) {
		return errors.New("not a valid url")
	}

	us.urls[token] = url

	return nil
}

func (us URLShortenerService) GetOriginalURL(shortURL string) (string, error) {
	if v, ok := us.urls[shortURL]; ok {
		return v, nil
	}

	msg := fmt.Sprintf("short url %s was never created", shortURL)

	return "", errors.New(msg)
}
