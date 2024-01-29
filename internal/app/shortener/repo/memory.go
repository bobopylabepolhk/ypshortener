package repo

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

	return "", errShortUrlDoesNotExist(shortURL)
}

func (repo *URLShortenerRepoMemory) SaveURLBatch(batch []URLBatch) error {
	for _, item := range batch {
		repo.urls[item.ShortURL] = item.OgURL
	}

	return nil
}

func newURLShortenerRepoMemory() *URLShortenerRepoMemory {
	return &URLShortenerRepoMemory{urls: make(map[string]string)}
}
