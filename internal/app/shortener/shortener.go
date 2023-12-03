package shortener

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	urlutils "github.com/bobopylabepolhk/ypshortener/pkg"
)

var urls = map[string]string{}

func GetShortUrlToken(l int) string {
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))

	result := make([]byte, l)

	for i := 0; i < l; i++ {
		result[i] = urlutils.ValidUrlSymbols[rand.Intn(len(urlutils.ValidUrlSymbols))]
	}

	return string(result)
}

func SaveShortUrl(url string, shortUrl string) error {
	if !urlutils.ValidateUrl(url) {
		return errors.New("not a valid url")
	}

	urls[shortUrl] = url

	return nil
}

func GetOriginalUrl(shortUrl string) (string, error) {
	if v, ok := urls[shortUrl]; ok {
		return v, nil
	}

	msg := fmt.Sprintf("short url %s was never created", shortUrl)

	return "", errors.New(msg)
}
