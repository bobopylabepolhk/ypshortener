package repo

import "database/sql"

type URLShortenerRepoPostgres struct {
	db *sql.DB
}

func (repo *URLShortenerRepoPostgres) CreateShortURL(token string, ogURL string) error {
	_, err := repo.db.Exec(
		"INSERT INTO url (og_url, short_url) VALUES ($1, $2)",
		ogURL,
		token,
	)
	return err
}

func (repo *URLShortenerRepoPostgres) GetOgURL(shortURL string) (string, error) {
	var ogURL string
	row := repo.db.QueryRow("SELECT og_url FROM url WHERE short_url = $1", shortURL)
	err := row.Scan(&ogURL)

	if err != nil {
		return "", errShortUrlDoesNotExist(shortURL)
	}

	return ogURL, nil
}

func newURLShortenerRepoPostgres(db *sql.DB) *URLShortenerRepoPostgres {
	return &URLShortenerRepoPostgres{db: db}
}
