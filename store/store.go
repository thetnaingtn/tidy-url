package store

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

//go:embed sql/schema.sql
var schema string

func NewStore(db *sqlx.DB) Store {
	return Store{
		db: db,
	}
}

func (s Store) Create(p Payload) (string, error) {

	id := uuid.New()
	urlPath := base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(id.ID())))

	t := TidyUrl{
		Id:       id.String(),
		LongURL:  p.LongURL,
		ShortURL: urlPath,
	}

	tx := s.db.MustBegin()
	tx.NamedExec(`INSERT INTO tidyurl (id,long_url,short_url) VALUES (:id,:long_url,:short_url)`, &t)
	tx.Commit()

	shortURL := fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), urlPath)
	return shortURL, nil
}
