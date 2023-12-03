package shortener

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	urlutils "github.com/bobopylabepolhk/ypshortener/pkg"
)

var urls = map[string]string{}

func GetShortURLToken(l int) string {
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))

	result := make([]byte, l)

	for i := 0; i < l; i++ {
		result[i] = urlutils.ValidURLSymbols[rand.Intn(len(urlutils.ValidURLSymbols))]
	}

	return string(result)
}

func SaveShortURL(url string, shortURL string) error {
	if !urlutils.ValidateURL(url) {
		return errors.New("not a valid url")
	}

	urls[shortURL] = url

	return nil
}

func GetOriginalURL(shortURL string) (string, error) {
	if v, ok := urls[shortURL]; ok {
		return v, nil
	}

	msg := fmt.Sprintf("short url %s was never created", shortURL)

	return "", errors.New(msg)
}
