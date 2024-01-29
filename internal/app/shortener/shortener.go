package shortener

import (
	"errors"

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

func (us URLShortenerService) SaveShortURL(url string, token string) error {
	if !urlutils.ValidateURL(url) {
		return errors.New("not a valid url")
	}

	return us.repo.CreateShortURL(token, url)
}

func (us URLShortenerService) GetOriginalURL(shortURL string) (string, error) {
	return us.repo.GetOgURL(shortURL)
}

func (us URLShortenerService) SaveURLBatch(batch []ShortenBatchRequestDTO) ([]ShortenBatchResponseDTO, error) {
	data := make([]repo.URLBatch, 0)
	res := make([]ShortenBatchResponseDTO, 0)
	for _, item := range batch {
		token := urlutils.GetShortURLToken()
		data = append(data, repo.URLBatch{ShortURL: token, OgURL: item.OgURL})
		res = append(res, ShortenBatchResponseDTO{ShortURL: token, CorrelationId: item.CorrelationId})
	}

	err := us.repo.SaveURLBatch(data)
	return res, err
}
