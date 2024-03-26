package repo

import (
	"context"
	"fmt"
)

type URLShortenerRepoMemory struct {
	urls map[string]string
}

func (repo *URLShortenerRepoMemory) CreateShortURL(_ context.Context, token string, ogURL string) error {
	repo.urls[token] = ogURL
	return nil
}

func (repo *URLShortenerRepoMemory) GetOgURL(_ context.Context, shortURL string) (string, error) {
	if v, ok := repo.urls[shortURL]; ok {
		return v, nil
	}

	return "", fmt.Errorf("memory.GetOgURL: %w", errShortURLDoesNotExist(shortURL))
}

func (repo *URLShortenerRepoMemory) SaveURLBatch(_ context.Context, batch []URLBatch) error {
	for _, item := range batch {
		repo.urls[item.ShortURL] = item.OgURL
	}

	return nil
}

func (repo *URLShortenerRepoMemory) FindTokenByOgURL(_ context.Context, ogURL string) (string, error) {
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
