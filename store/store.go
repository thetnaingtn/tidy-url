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

func (s Store) GetLongURL(encodedPath string) (TidyUrl, error) {
	var tidyurl TidyUrl
	err := s.db.Get(&tidyurl, "SELECT * FROM tidyurl WHERE short_url=$1", encodedPath)

	if err != nil {
		return TidyUrl{}, err
	}

	return tidyurl, nil

}
