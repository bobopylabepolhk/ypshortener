package shortener

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/bobopylabepolhk/ypshortener/config"
	urlutils "github.com/bobopylabepolhk/ypshortener/pkg"
)

var urls = map[string]string{}

type URLShortenerService struct {
	urls     map[string]string
	tokenLen int
}

func NewURLShortenerService() *URLShortenerService {
	return &URLShortenerService{
		urls:     make(map[string]string),
		tokenLen: config.MinTokenLength,
	}
}

func (us URLShortenerService) GetShortURLToken() string {
	const tokenChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))

	result := make([]byte, us.tokenLen)

	for i := 0; i < us.tokenLen; i++ {
		result[i] = tokenChars[rand.Intn(len(tokenChars))]
	}

	return string(result)
}

func (us URLShortenerService) SaveShortURL(url string, token string) error {
	if !urlutils.ValidateURL(url) {
		return errors.New("not a valid url")
	}

	urls[token] = url

	return nil
}

func (us URLShortenerService) GetOriginalURL(shortURL string) (string, error) {
	if v, ok := urls[shortURL]; ok {
		return v, nil
	}

	msg := fmt.Sprintf("short url %s was never created", shortURL)

	return "", errors.New(msg)
}
