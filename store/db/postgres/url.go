package postgres

import (
	"context"

	"github.com/thetnaingtn/tidy-url/store"
)

var _ store.Driver = (*DB)(nil)

func (driver *DB) Create(ctx context.Context, t *store.TidyUrl) error {
	stmt := `
		INSERT INTO tidy_url (long_url, encoded_str)
		VALUES ($1, $2)
		RETURNING id, created_at
	`

	args := []any{t.LongUrl, t.EncodedStr}

	return driver.db.QueryRowContext(ctx, stmt, args...).Scan(
		&t.Id,
		&t.CreatedAt,
	)
}

func (driver *DB) FindTidyUrl(ctx context.Context, filter *store.Filters) (*store.TidyUrl, error) {
	stmt := `
		SELECT id, long_url, encoded_str, created_at
		FROM tidy_url
		WHERE encoded_str = $1 OR long_url = $2
		LIMIT 1
	`

	args := []any{filter.EncodedStr, filter.LongUrl}

	tidyUrl := &store.TidyUrl{}
	err := driver.db.QueryRowContext(ctx, stmt, args...).Scan(
		&tidyUrl.Id,
		&tidyUrl.LongUrl,
		&tidyUrl.EncodedStr,
		&tidyUrl.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return tidyUrl, nil
}
