package store

import (
	"context"

	"github.com/lithammer/shortuuid/v4"
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

func (s *Store) Create(ctx context.Context, url string) (*TidyUrl, error) {
	encodedStr := shortuuid.NewWithNamespace(url)
	tidyUrl := &TidyUrl{
		LongUrl:    url,
		EncodedStr: encodedStr,
	}

	if err := s.driver.Create(ctx, tidyUrl); err != nil {
		return nil, err
	}

	return tidyUrl, nil
}
