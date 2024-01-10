package shortener

import (
	"errors"
	"fmt"

	"github.com/bobopylabepolhk/ypshortener/config"
	"github.com/bobopylabepolhk/ypshortener/pkg/jsonreader"
)

type (
	URLShortenerRow struct {
		id       int
		shortURL string
		ogURL    string
	}

	JSONDbReader interface {
		WriteRow(data interface{}) error
		FindByKey(key string) ([]byte, error)
	}

	URLShortenerRepo struct {
		useJSONReader bool
		reader        JSONDbReader
		urls          map[string]string
	}
)

func (repo URLShortenerRepo) CreateShortURL(token string, ogURL string) {
	repo.urls[token] = ogURL
}

func (repo URLShortenerRepo) GetOgURL(shortURL string) (string, error) {
	if v, ok := repo.urls[shortURL]; ok {
		return v, nil
	}

	msg := fmt.Sprintf("short url %s was never created", shortURL)

	return "", errors.New(msg)
}

func NewURLShortenerRepo() *URLShortenerRepo {
	JSONreader, err := jsonreader.NewJSONReader(config.Cfg.UrlStoragePath)
	useJSONReader := err != nil || config.Cfg.UrlStoragePath == ""

	return &URLShortenerRepo{
		useJSONReader: useJSONReader,
		reader:        JSONreader,
		urls:          make(map[string]string),
	}
}
