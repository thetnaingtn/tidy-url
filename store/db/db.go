package db

import (
	"github.com/thetnaingtn/tidy-url/internal/config"
	"github.com/thetnaingtn/tidy-url/store"
	"github.com/thetnaingtn/tidy-url/store/db/postgres"
)

func NewDBDriver(config *config.Config) (store.Driver, error) {
	db, err := postgres.NewDB(config)
	if err != nil {
		return nil, err
	}

	return db, nil
}
