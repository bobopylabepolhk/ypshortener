package repo

import (
	"context"
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

	URLBatch struct {
		ShortURL string
		OgURL    string
	}

	URLShortenerRepository interface {
		CreateShortURL(ctx context.Context, token string, ogURL string) error
		GetOgURL(ctx context.Context, shortURL string) (string, error)
		SaveURLBatch(ctx context.Context, batch []URLBatch) error
		FindTokenByOgURL(ctx context.Context, ogURL string) (string, error)
	}
)

const (
	RepoMemory URLShortenerRepositoryKind = iota
	RepoPostgres
	RepoWithJSONReader
)

func errShortURLDoesNotExist(shortURL string) error {
	return fmt.Errorf("short url %s was never created", shortURL)
}

func errOgURLNotFound(ogURL string) error {
	return fmt.Errorf("original url %s was not found", ogURL)
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
	case RepoWithJSONReader:
		return newURLShortenerRepoWithReader(repoConfig.storagePath)
	default:
		return nil, fmt.Errorf("invalid URLShortenerStorageRepository kind")
	}
}

func WithJSONReader() option {
	return func(repoCfg *URLShortenerRepositoryConfig) {
		if config.Cfg.URLStoragePath != "" {
			repoCfg.kind = RepoWithJSONReader
			repoCfg.storagePath = config.Cfg.URLStoragePath
		} else {
			repoCfg.kind = RepoMemory
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

	return WithJSONReader()
}
