package repo

import (
	"context"
	"fmt"
)

type URLShortenerRepoMemory struct {
	urls         map[string]string
	urlsByUserID map[string][]URLBatch
}

func (repo *URLShortenerRepoMemory) CreateShortURL(_ context.Context, token string, ogURL string, userID string) error {
	repo.urls[token] = ogURL
	repo.urlsByUserID[userID] = append(repo.urlsByUserID[userID], URLBatch{ShortURL: token, OgURL: ogURL})
	return nil
}

func (repo *URLShortenerRepoMemory) GetOgURL(_ context.Context, shortURL string) (string, error) {
	if v, ok := repo.urls[shortURL]; ok {
		return v, nil
	}

	return "", fmt.Errorf("memory.GetOgURL: %w", errShortURLDoesNotExist(shortURL))
}

func (repo *URLShortenerRepoMemory) SaveURLBatch(_ context.Context, batch []URLBatch, userID string) error {
	newUserRecords := []URLBatch{}
	for idx, item := range batch {
		repo.urls[item.ShortURL] = item.OgURL
		newUserRecords[idx] = URLBatch{ShortURL: item.ShortURL, OgURL: item.OgURL}
	}
	repo.urlsByUserID[userID] = append(repo.urlsByUserID[userID], newUserRecords...)

	return nil
}

func (repo *URLShortenerRepoMemory) FindTokenByOgURL(_ context.Context, ogURL string) (string, error) {
	for short, og := range repo.urls {
		if og == ogURL {
			return short, nil
		}
	}

	return "", fmt.Errorf("memory.FindTokenByOgURLs: %w", errOgURLNotFound(ogURL))
}

func (repo *URLShortenerRepoMemory) GetURLsByUser(_ context.Context, userID string) ([]URLBatch, error) {
	return repo.urlsByUserID[userID], nil
}

func (repo *URLShortenerRepoMemory) DeleteURLs(ctx context.Context, tokens []string, userID string) error {
	for _, k := range repo.urls {
		delete(repo.urls, k)
		delete(repo.urlsByUserID, k)
	}
	return nil
}

func newURLShortenerRepoMemory() *URLShortenerRepoMemory {
	return &URLShortenerRepoMemory{urls: make(map[string]string), urlsByUserID: make(map[string][]URLBatch)}
}
