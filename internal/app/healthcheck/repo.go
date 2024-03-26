package healthcheck

import (
	"database/sql"
)

type HealthcheckRepo struct {
	db *sql.DB
}

func (repo *HealthcheckRepo) Ping() error {
	return repo.db.Ping()
}

func NewHealthcheckRepo(db *sql.DB) *HealthcheckRepo {
	return &HealthcheckRepo{
		db: db,
	}
}
