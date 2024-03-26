package repo

import (
	"fmt"

	"github.com/bobopylabepolhk/ypshortener/pkg/jsonreader"
)

type (
	URLShortenerRow struct {
		ShortURL string `json:"short_url"`
		OgURL    string `json:"original_url"`
	}

	URLShortenerRepoWithJSONReader struct {
		jsonReader JSONDbReader
		repo       URLShortenerRepoMemory
	}

	JSONDbReader interface {
		WriteRow(data interface{}) error
		InitFromFile() ([]map[string]interface{}, error)
	}
)

func (repoWithReader *URLShortenerRepoWithJSONReader) CreateShortURL(token string, ogURL string) error {
	data := URLShortenerRow{ShortURL: token, OgURL: ogURL}
	err := repoWithReader.repo.CreateShortURL(token, ogURL)

	if err != nil {
		return fmt.Errorf("jsonReader.CreateShortURL: %w", err)
	}

	return repoWithReader.jsonReader.WriteRow(data)
}

func (repoWithReader *URLShortenerRepoWithJSONReader) GetOgURL(shortURL string) (string, error) {
	return repoWithReader.repo.GetOgURL(shortURL)
}

func (repoWithReader *URLShortenerRepoWithJSONReader) FindTokenByOgURL(ogURL string) (string, error) {
	return repoWithReader.repo.FindTokenByOgURL(ogURL)
}

func (repoWithReader *URLShortenerRepoWithJSONReader) SaveURLBatch(batch []URLBatch) error {
	for _, item := range batch {
		err := repoWithReader.jsonReader.WriteRow(item)

		if err != nil {
			return fmt.Errorf("jsonReader.SaveURLBatch: %w", err)
		}
	}

	return repoWithReader.repo.SaveURLBatch(batch)
}

func newURLShortenerRepoWithReader(storagePath string) (*URLShortenerRepoWithJSONReader, error) {
	JSONReader, err := jsonreader.NewJSONReader(storagePath)
	if err != nil {
		return nil, fmt.Errorf("jsonReader: %w", err)
	}

	urls := map[string]string{}
	json, err := JSONReader.InitFromFile()

	if err != nil {
		return nil, fmt.Errorf("jsonReader: %w", err)
	}

	for _, item := range json {
		key := item["short_url"].(string)
		v := item["original_url"].(string)

		urls[key] = v
	}

	repo := URLShortenerRepoMemory{urls: urls}
	return &URLShortenerRepoWithJSONReader{repo: repo, jsonReader: JSONReader}, nil
}
