package postgres

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	"github.com/thetnaingtn/tidy-url/internal/config"
)

var ErrNoDSN = errors.New("no DSN provided")

type DB struct {
	db     *sql.DB
	config *config.Config
}

func NewDB(cfg *config.Config) (*DB, error) {
	if cfg.DSN == "" {
		return nil, ErrNoDSN
	}

	db, err := sql.Open("postgres", cfg.DSN)

	if err != nil {
		return nil, err
	}

	return &DB{
		db:     db,
		config: cfg,
	}, nil
}
