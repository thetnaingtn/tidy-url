package core

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/thetnaingtn/tidy-url/store"
)

type Core struct {
	store store.Store
}

// TODO: FIX LATER
func NewCore(db *sqlx.DB) Core {
	store, _ := store.NewStore(nil)
	return Core{
		store: *store,
	}
}

func (c Core) GenerateTidyUrl(p store.Payload) (string, error) {
	tidyurl, _ := c.store.GetRecordByLongURL(p.LongURL)
	if tidyurl.ShortURL != "" {
		return fmt.Sprintf("%s/expand/%s", os.Getenv("BASE_URL"), tidyurl.ShortURL), nil
	}

	urlPath, err := c.store.Create(p)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/expand/%s", os.Getenv("BASE_URL"), urlPath), nil
}

func (c Core) GetLongURL(encodedString string) (store.TidyUrl, error) {
	tidyurl, err := c.store.GetRecordByShortURL(encodedString)
	if err != nil {
		return store.TidyUrl{}, err
	}

	return tidyurl, nil
}
