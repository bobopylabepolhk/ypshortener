package repo

import (
	"database/sql"
	"errors"
)

type URLShortenerRepoPostgres struct {
	db *sql.DB
}

var ErrDuplicateURL = errors.New("shortURL already exists for this ogURL")

func (repo *URLShortenerRepoPostgres) CreateShortURL(token string, ogURL string) error {
	res, err := repo.db.Exec(
		"INSERT INTO url (og_url, short_url) VALUES ($1, $2) ON CONFLICT DO NOTHING",
		ogURL,
		token,
	)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrDuplicateURL
	}

	return nil
}

func (repo *URLShortenerRepoPostgres) GetOgURL(shortURL string) (string, error) {
	var ogURL string
	row := repo.db.QueryRow("SELECT og_url FROM url WHERE short_url = $1", shortURL)
	err := row.Scan(&ogURL)

	if err != nil {
		return "", errShortURLDoesNotExist(shortURL)
	}

	return ogURL, nil
}

func (repo *URLShortenerRepoPostgres) SaveURLBatch(batch []URLBatch) error {
	t, err := repo.db.Begin()

	if err != nil {
		return err
	}

	defer t.Rollback()

	stmt, err := t.Prepare("INSERT INTO url (og_url, short_url) VALUES ($1, $2) ON CONFLICT (og_url) DO NOTHING")
	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, item := range batch {
		_, err := stmt.Exec(item.OgURL, item.ShortURL)

		if err != nil {
			return err
		}
	}

	return t.Commit()
}

func (repo *URLShortenerRepoPostgres) FindTokenByOgURL(ogURL string) (string, error) {
	var res string
	row := repo.db.QueryRow("SELECT short_url FROM url WHERE og_url = $1", ogURL)
	err := row.Scan(&res)

	if err != nil {
		return "", errOgURLNotFound(ogURL)
	}

	return res, nil
}

func newURLShortenerRepoPostgres(db *sql.DB) *URLShortenerRepoPostgres {
	return &URLShortenerRepoPostgres{db: db}
}
