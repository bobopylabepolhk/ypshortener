package db

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/bobopylabepolhk/ypshortener/config"
)

func New() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.Cfg.PostgresDSN)

	if err != nil {
		return nil, err
	}

	return db, nil
}
