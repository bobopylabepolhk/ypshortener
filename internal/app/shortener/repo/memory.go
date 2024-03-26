package repo

import "fmt"

type URLShortenerRepoMemory struct {
	urls map[string]string
}

func (repo *URLShortenerRepoMemory) CreateShortURL(token string, ogURL string) error {
	repo.urls[token] = ogURL
	return nil
}

func (repo *URLShortenerRepoMemory) GetOgURL(shortURL string) (string, error) {
	if v, ok := repo.urls[shortURL]; ok {
		return v, nil
	}

	return "", fmt.Errorf("memory.GetOgURL: %w", errShortURLDoesNotExist(shortURL))
}

func (repo *URLShortenerRepoMemory) SaveURLBatch(batch []URLBatch) error {
	for _, item := range batch {
		repo.urls[item.ShortURL] = item.OgURL
	}

	return nil
}

func (repo *URLShortenerRepoMemory) FindTokenByOgURL(ogURL string) (string, error) {
	for short, og := range repo.urls {
		if og == ogURL {
			return short, nil
		}
	}

	return "", fmt.Errorf("memory.GetOgURL: %w", errOgURLNotFound(ogURL))
}

func newURLShortenerRepoMemory() *URLShortenerRepoMemory {
	return &URLShortenerRepoMemory{urls: make(map[string]string)}
}
