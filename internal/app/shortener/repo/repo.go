package repo

import (
	"database/sql"
	"fmt"

	"github.com/bobopylabepolhk/ypshortener/config"
)

type (
	URLShortenerRepositoryKind int

	URLShortenerRepositoryConfig struct {
		kind        URLShortenerRepositoryKind
		db          *sql.DB
		storagePath string
	}

	option func(*URLShortenerRepositoryConfig)

	URLShortenerRepository interface {
		CreateShortURL(token string, ogURL string) error
		GetOgURL(shortURL string) (string, error)
	}
)

const (
	RepoMemory URLShortenerRepositoryKind = iota
	RepoPostgres
	RepoWithJsonReader
)

func errShortUrlDoesNotExist(shortURL string) error {
	return fmt.Errorf("short url %s was never created", shortURL)
}

func NewURLShortenerRepo(options ...option) (URLShortenerRepository, error) {
	repoConfig := &URLShortenerRepositoryConfig{
		kind: RepoMemory,
	}

	for _, opt := range options {
		opt(repoConfig)
	}

	switch repoConfig.kind {
	case RepoMemory:
		return newURLShortenerRepoMemory(), nil
	case RepoPostgres:
		return newURLShortenerRepoPostgres(repoConfig.db), nil
	case RepoWithJsonReader:
		return newURLShortenerRepoWithReader(repoConfig.storagePath)
	default:
		return nil, fmt.Errorf("invalid URLShortenerStorageRepository kind")
	}
}

func WithJsonReader() option {
	return func(repoCfg *URLShortenerRepositoryConfig) {
		if config.Cfg.URLStoragePath != "" {
			repoCfg.kind = RepoWithJsonReader
			repoCfg.storagePath = config.Cfg.URLStoragePath
		}
	}
}

func WithPostgres(db *sql.DB) option {
	if config.Cfg.PostgresDSN != "" {
		return func(repoCfg *URLShortenerRepositoryConfig) {
			repoCfg.kind = RepoPostgres
			repoCfg.db = db
		}
	}

	return WithJsonReader()
}
