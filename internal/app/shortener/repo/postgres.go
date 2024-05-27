package repo

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/bobopylabepolhk/ypshortener/pkg/logger"
)

type URLShortenerRepoPostgres struct {
	db *sql.DB
}

func (repo *URLShortenerRepoPostgres) CreateShortURL(ctx context.Context, token string, ogURL string, userID string) error {
	res, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO url (og_url, short_url, user_id, is_deleted) VALUES ($1, $2, $3, false) ON CONFLICT DO NOTHING",
		ogURL,
		token,
		userID,
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
	var isDeleted bool
	row := repo.db.QueryRowContext(ctx, "SELECT og_url, is_deleted FROM url WHERE short_url = $1", shortURL)
	err := row.Scan(&ogURL, &isDeleted)

	if err != nil {
		return "", fmt.Errorf("postgres.GetOgURL: %w", errShortURLDoesNotExist(shortURL))
	}

	if isDeleted {
		return "", ErrURLIsDeleted
	}

	return ogURL, nil
}

func (repo *URLShortenerRepoPostgres) SaveURLBatch(ctx context.Context, batch []URLBatch, userID string) error {
	t, err := repo.db.BeginTx(ctx, nil)

	if err != nil {
		return fmt.Errorf("postgres.SaveURLBatch: %w", err)
	}

	defer t.Rollback()

	stmt, err := t.Prepare("INSERT INTO url (og_url, short_url, user_id, is_deleted) VALUES ($1, $2, $3, false) ON CONFLICT (og_url) DO NOTHING")
	if err != nil {
		return fmt.Errorf("postgres.SaveURLBatch: %w", err)
	}

	defer stmt.Close()

	for _, item := range batch {
		_, err := stmt.ExecContext(ctx, item.OgURL, item.ShortURL, userID)

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

func (repo *URLShortenerRepoPostgres) GetURLsByUser(ctx context.Context, userID string) ([]URLBatch, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT short_url, og_url FROM url WHERE user_id = $1 AND is_deleted != true", userID)
	if err != nil {
		return nil, fmt.Errorf("postgres.GetURLsByUser: %w", err)
	}

	defer rows.Close()

	res := []URLBatch{}
	for rows.Next() {
		item := URLBatch{}
		if rows.Err() != nil {
			return res, fmt.Errorf("postgres.GetURLsByUser: %w", rows.Err())
		}

		if err = rows.Scan(&item.ShortURL, &item.OgURL); err != nil {
			return res, fmt.Errorf("postgres.GetURLsByUser: %w", err)
		}
		res = append(res, item)
	}

	return res, nil
}

func (repo *URLShortenerRepoPostgres) DeleteURLs(ctx context.Context, tokens []string, userID string) error {
	t, err := repo.db.BeginTx(ctx, nil)

	if err != nil {
		return fmt.Errorf("postgres.DeleteURLs: %w", err)
	}

	defer t.Rollback()

	stmt, err := t.Prepare("UPDATE url SET is_deleted = true WHERE user_id=$1 AND short_url=$2")
	if err != nil {
		return fmt.Errorf("postgres.DeleteURLs: %w", err)
	}

	defer stmt.Close()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg := &sync.WaitGroup{}
	for _, token := range tokens {
		wg.Add(1)
		go func(token string) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
			}

			_, err := stmt.ExecContext(ctx, userID, token)
			if err != nil {
				logger.Error(fmt.Sprintf("postgres.DeleteURLs failed to delete %v: %v", token, err.Error()))
				cancel()
			}
		}(token)
	}

	wg.Wait()
	err = t.Commit()
	if err != nil {
		return fmt.Errorf("postgres.DeleteURLs failed to commit: %w", err)
	}

	return nil
}

func newURLShortenerRepoPostgres(db *sql.DB) *URLShortenerRepoPostgres {
	return &URLShortenerRepoPostgres{db: db}
}
