package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type URLShortenerRepoPostgres struct {
	db *sql.DB
}

var ErrDuplicateURL = errors.New("shortURL already exists for this ogURL")

func (repo *URLShortenerRepoPostgres) CreateShortURL(ctx context.Context, token string, ogURL string) error {
	res, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO url (og_url, short_url) VALUES ($1, $2) ON CONFLICT DO NOTHING",
		ogURL,
		token,
	)

	if err != nil {
		return fmt.Errorf("postgres.CreateShortURL: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("postgres.CreateShortURL: %w", err)
	}

	if rows == 0 {
		return ErrDuplicateURL
	}

	return nil
}

func (repo *URLShortenerRepoPostgres) GetOgURL(ctx context.Context, shortURL string) (string, error) {
	var ogURL string
	row := repo.db.QueryRowContext(ctx, "SELECT og_url FROM url WHERE short_url = $1", shortURL)
	err := row.Scan(&ogURL)

	if err != nil {
		return "", fmt.Errorf("postgres.GetOgURL: %w", errShortURLDoesNotExist(shortURL))
	}

	return ogURL, nil
}

func (repo *URLShortenerRepoPostgres) SaveURLBatch(ctx context.Context, batch []URLBatch) error {
	t, err := repo.db.BeginTx(ctx, nil)

	if err != nil {
		return fmt.Errorf("postgres.SaveURLBatch: %w", err)
	}

	defer t.Rollback()

	stmt, err := t.Prepare("INSERT INTO url (og_url, short_url) VALUES ($1, $2) ON CONFLICT (og_url) DO NOTHING")
	if err != nil {
		return fmt.Errorf("postgres.SaveURLBatch: %w", err)
	}

	defer stmt.Close()

	for _, item := range batch {
		_, err := stmt.ExecContext(ctx, item.OgURL, item.ShortURL)

		if err != nil {
			return fmt.Errorf("postgres.SaveURLBatch: %w", err)
		}
	}

	return t.Commit()
}

func (repo *URLShortenerRepoPostgres) FindTokenByOgURL(ctx context.Context, ogURL string) (string, error) {
	var res string
	row := repo.db.QueryRowContext(ctx, "SELECT short_url FROM url WHERE og_url = $1", ogURL)
	err := row.Scan(&res)

	if err != nil {
		return "", fmt.Errorf("postgres.FindTokenByOgURL: %w", errOgURLNotFound(ogURL))
	}

	return res, nil
}

func newURLShortenerRepoPostgres(db *sql.DB) *URLShortenerRepoPostgres {
	return &URLShortenerRepoPostgres{db: db}
}
