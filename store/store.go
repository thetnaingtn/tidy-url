package store

import (
	"github.com/thetnaingtn/tidy-url/internal/config"
)

type Store struct {
	config *config.Config
	driver Driver
}

func NewStore(config *config.Config, driver Driver) *Store {
	return &Store{
		config: config,
		driver: driver,
	}
}
