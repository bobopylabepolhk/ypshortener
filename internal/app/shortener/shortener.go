package shortener

import (
	"context"
	"errors"
	"fmt"

	"github.com/bobopylabepolhk/ypshortener/config"
	"github.com/bobopylabepolhk/ypshortener/internal/app/shortener/repo"
	"github.com/bobopylabepolhk/ypshortener/pkg/urlutils"
)

type (
	URLShortenerService struct {
		repo repo.URLShortenerRepository
	}
)

func NewURLShortenerService(repo repo.URLShortenerRepository) *URLShortenerService {
	return &URLShortenerService{
		repo: repo,
	}
}

func (us URLShortenerService) SaveShortURL(ctx context.Context, url string, token string, userID string) (string, error) {
	if !urlutils.ValidateURL(url) {
		return "", errors.New("not a valid url")
	}

	shortURL := fmt.Sprintf("%s/%s", config.Cfg.BaseURL, token)
	err := us.repo.CreateShortURL(ctx, token, url, userID)

	return shortURL, err
}

func (us URLShortenerService) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	return us.repo.GetOgURL(ctx, shortURL)
}

func (us URLShortenerService) GetExistingShortURL(ctx context.Context, ogURL string) (string, error) {
	token, err := us.repo.FindTokenByOgURL(ctx, ogURL)
	return fmt.Sprintf("%s/%s", config.Cfg.BaseURL, token), err
}

func (us URLShortenerService) SaveURLBatch(ctx context.Context, batch []ShortenBatchRequestDTO, userID string) ([]ShortenBatchResponseDTO, error) {
	data := make([]repo.URLBatch, 0)
	res := make([]ShortenBatchResponseDTO, 0)
	for _, item := range batch {
		token := urlutils.CreateRandomToken(6)
		data = append(data, repo.URLBatch{ShortURL: token, OgURL: item.OgURL})
		shortURL := fmt.Sprintf("%s/%s", config.Cfg.BaseURL, token)
		res = append(res, ShortenBatchResponseDTO{ShortURL: shortURL, CorrelationID: item.CorrelationID})
	}

	err := us.repo.SaveURLBatch(ctx, data, userID)
	return res, err
}

func (us URLShortenerService) GetUserURLs(ctx context.Context, userID string) ([]repo.URLBatch, error) {
	res, err := us.repo.GetURLsByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	for idx, item := range res {
		res[idx].ShortURL = fmt.Sprintf("%s/%s", config.Cfg.BaseURL, item.ShortURL)
	}

	return res, nil
}

func (us URLShortenerService) DeleteURLs(ctx context.Context, tokens []string, userID string) error {
	return us.repo.DeleteURLs(ctx, tokens, userID)
}
