package shortener

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	urlutils "github.com/bobopylabepolhk/ypshortener/pkg"
)

var urls = map[string]string{}

type URLShortener struct {
	urls     map[string]string
	tokenLen int
}

func NewURLShortener(l int) *URLShortener {
	return &URLShortener{urls: make(map[string]string), tokenLen: l}
}

func (us URLShortener) GetShortURLToken() string {
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))

	result := make([]byte, us.tokenLen)

	for i := 0; i < us.tokenLen; i++ {
		result[i] = urlutils.ValidURLSymbols[rand.Intn(len(urlutils.ValidURLSymbols))]
	}

	return string(result)
}

func (us URLShortener) SaveShortURL(url string, token string) error {
	if !urlutils.ValidateURL(url) {
		return errors.New("not a valid url")
	}

	urls[token] = url

	return nil
}

func (us URLShortener) GetOriginalURL(shortURL string) (string, error) {
	if v, ok := urls[shortURL]; ok {
		return v, nil
	}

	msg := fmt.Sprintf("short url %s was never created", shortURL)

	return "", errors.New(msg)
}
