package store

import (
	"context"
	"time"

	nanoid "github.com/matoous/go-nanoid/v2"
)

type Filters struct {
	EncodedStr string
	LongUrl    string
}

type TidyUrl struct {
	Id         string
	LongUrl    string
	EncodedStr string
	CreatedAt  time.Time
}

func (s *Store) Create(ctx context.Context, url string) (*TidyUrl, error) {
	encodedStr, err := nanoid.New(10)
	if err != nil {
		return nil, err
	}

	tidyUrl := &TidyUrl{
		LongUrl:    url,
		EncodedStr: encodedStr,
	}

	if err := s.driver.Create(ctx, tidyUrl); err != nil {
		return nil, err
	}

	return tidyUrl, nil
}

func (s *Store) GetTidyUrlByOriginalUrl(ctx context.Context, longUrl string) (*TidyUrl, error) {
	tidyUrl, err := s.driver.FindTidyUrl(ctx, &Filters{LongUrl: longUrl})

	if err != nil {
		return nil, err
	}

	return tidyUrl, nil
}

func (s *Store) GetTidyUrlByEncodedStr(ctx context.Context, encodedStr string) (*TidyUrl, error) {
	tidyUrl, err := s.driver.FindTidyUrl(ctx, &Filters{EncodedStr: encodedStr})

	if err != nil {
		return nil, err
	}

	return tidyUrl, nil
}
