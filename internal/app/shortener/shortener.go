package shortener

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

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

func (us URLShortenerService) GetShortURLToken() string {
	const tokenChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))

	tokenLen := 6
	result := make([]byte, tokenLen)

	for i := 0; i < tokenLen; i++ {
		result[i] = tokenChars[rand.Intn(len(tokenChars))]
	}

	return string(result)
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
