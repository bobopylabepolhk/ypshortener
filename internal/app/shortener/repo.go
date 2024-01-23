package shortener

import (
	"errors"
	"fmt"

	"github.com/bobopylabepolhk/ypshortener/config"
	jsonreader "github.com/bobopylabepolhk/ypshortener/pkg/jsonreader"
)

type (
	URLShortenerRow struct {
		ShortURL string `json:"short_url"`
		OgURL    string `json:"original_url"`
	}

	JSONDbReader interface {
		WriteRow(data interface{}) error
		InitFromFile() ([]map[string]interface{}, error)
	}

	URLShortenerRepo struct {
		useJSONReader bool
		jsonReader    JSONDbReader
		urls          map[string]string
	}
)

func (repo *URLShortenerRepo) CreateShortURL(token string, ogURL string) {
	repo.urls[token] = ogURL
	if repo.useJSONReader {
		data := URLShortenerRow{ShortURL: token, OgURL: ogURL}
		repo.jsonReader.WriteRow(data)
	}
}

func (repo *URLShortenerRepo) GetOgURL(shortURL string) (string, error) {
	if v, ok := repo.urls[shortURL]; ok {
		return v, nil
	}

	msg := fmt.Sprintf("short url %s was never created", shortURL)

	return "", errors.New(msg)
}

func NewURLShortenerRepo() *URLShortenerRepo {
	JSONReader, err := jsonreader.NewJSONReader(config.Cfg.URLStoragePath)
	urls := map[string]string{}
	useJSONReader := err == nil && config.Cfg.URLStoragePath != ""

	if useJSONReader {
		json, err := JSONReader.InitFromFile()
		if err == nil {
			for _, item := range json {
				key := item["short_url"].(string)
				v := item["original_url"].(string)

				urls[key] = v
			}
		}
	}

	return &URLShortenerRepo{
		useJSONReader: useJSONReader,
		jsonReader:    JSONReader,
		urls:          urls,
	}
}
