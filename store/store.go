package store

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/thetnaingtn/tidy-url/internal/config"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(cfg *config.Config) (*Store, error) {
	db, err := sqlx.Connect("postgres", cfg.DSN)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

func (s Store) Create(p Payload) (string, error) {

	id := uuid.New()
	urlPath := base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(id.ID())))

	t := TidyUrl{
		Id:        id.String(),
		LongURL:   p.LongURL,
		ShortURL:  urlPath,
		CreatedAt: time.Now(),
	}

	tx := s.db.MustBegin()
	if _, err := tx.NamedExec(`INSERT INTO tidyurl (id,long_url,short_url,created_at) VALUES (:id,:long_url,:short_url,:created_at)`, &t); err != nil {
		return "", err
	}
	tx.Commit()

	return urlPath, nil
}

func (s Store) GetRecordByLongURL(longURL string) (TidyUrl, error) {
	var tidyurl TidyUrl

	err := s.db.Get(&tidyurl, "SELECT * FROM tidyurl WHERE long_url=$1", longURL)
	if err != nil {
		return TidyUrl{}, err
	}

	return tidyurl, nil
}

func (s Store) GetRecordByShortURL(encodedPath string) (TidyUrl, error) {
	var tidyurl TidyUrl
	err := s.db.Get(&tidyurl, "SELECT * FROM tidyurl WHERE short_url=$1", encodedPath)

	if err != nil {
		return TidyUrl{}, err
	}

	return tidyurl, nil
}
