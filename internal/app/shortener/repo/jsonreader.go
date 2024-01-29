package repo

import "github.com/bobopylabepolhk/ypshortener/pkg/jsonreader"

type (
	URLShortenerRow struct {
		ShortURL string `json:"short_url"`
		OgURL    string `json:"original_url"`
	}

	URLShortenerRepoWithJsonReader struct {
		jsonReader JSONDbReader
		repo       URLShortenerRepoMemory
	}

	JSONDbReader interface {
		WriteRow(data interface{}) error
		InitFromFile() ([]map[string]interface{}, error)
	}
)

func (repoWithReader *URLShortenerRepoWithJsonReader) CreateShortURL(token string, ogURL string) error {
	data := URLShortenerRow{ShortURL: token, OgURL: ogURL}
	err := repoWithReader.repo.CreateShortURL(token, ogURL)

	if err != nil {
		return err
	}

	return repoWithReader.jsonReader.WriteRow(data)
}

func (repoWithReader *URLShortenerRepoWithJsonReader) GetOgURL(shortURL string) (string, error) {
	return repoWithReader.repo.GetOgURL(shortURL)
}

func newURLShortenerRepoWithReader(storagePath string) (*URLShortenerRepoWithJsonReader, error) {
	JSONReader, err := jsonreader.NewJSONReader(storagePath)
	if err != nil {
		return nil, err
	}

	urls := map[string]string{}
	json, err := JSONReader.InitFromFile()

	if err != nil {
		return nil, err
	}

	for _, item := range json {
		key := item["short_url"].(string)
		v := item["original_url"].(string)

		urls[key] = v
	}

	repo := URLShortenerRepoMemory{urls: urls}
	return &URLShortenerRepoWithJsonReader{repo: repo, jsonReader: JSONReader}, nil
}
