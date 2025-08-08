package postgres

import (
	"database/sql"

	"github.com/thetnaingtn/tidy-url/internal/config"
)

type DB struct {
	db     *sql.DB
	config *config.Config
}

func NewDB(cfg *config.Config) (*DB, error) {
	db, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		return nil, err
	}

	return &DB{
		db:     db,
		config: cfg,
	}, nil
}
