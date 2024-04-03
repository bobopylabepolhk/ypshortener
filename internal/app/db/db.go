package db

import (
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"

	"github.com/bobopylabepolhk/ypshortener/config"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func New() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.Cfg.PostgresDSN)

	if err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}

	return db, nil
}

func Migrate(db *sql.DB) error {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFS,
		Root:       "migrations",
	}

	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)

	return fmt.Errorf("failed to run migrations: %w", err)
}
