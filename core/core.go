package core

import (
	"github.com/jmoiron/sqlx"
	"github.com/thetnaingtn/tidy-url/store"
)

type Core struct {
	store store.Store
}

func NewCore(db *sqlx.DB) Core {
	return Core{
		store: store.NewStore(db),
	}
}

func (c Core) GenerateTidyUrl(p store.Payload) (string, error) {
	id, err := c.store.Create(p)
	if err != nil {
		return "", err
	}

	return id, nil
}
