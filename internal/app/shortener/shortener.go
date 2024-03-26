package shortener

import (
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

func (us URLShortenerService) SaveShortURL(url string, token string) (string, error) {
	if !urlutils.ValidateURL(url) {
		return "", errors.New("not a valid url")
	}

	return fmt.Sprintf("%s/%s", config.Cfg.BaseURL, token), us.repo.CreateShortURL(token, url)
}

func (us URLShortenerService) GetOriginalURL(shortURL string) (string, error) {
	return us.repo.GetOgURL(shortURL)
}

func (us URLShortenerService) GetExistingShortURL(ogURL string) (string, error) {
	token, err := us.repo.FindTokenByOgURL(ogURL)
	return fmt.Sprintf("%s/%s", config.Cfg.BaseURL, token), err
}

func (us URLShortenerService) SaveURLBatch(batch []ShortenBatchRequestDTO) ([]ShortenBatchResponseDTO, error) {
	data := make([]repo.URLBatch, 0)
	res := make([]ShortenBatchResponseDTO, 0)
	for _, item := range batch {
		token := urlutils.GetShortURLToken()
		data = append(data, repo.URLBatch{ShortURL: token, OgURL: item.OgURL})
		shortURL := fmt.Sprintf("%s/%s", config.Cfg.BaseURL, token)
		res = append(res, ShortenBatchResponseDTO{ShortURL: shortURL, CorrelationID: item.CorrelationID})
	}

	err := us.repo.SaveURLBatch(data)
	return res, err
}
