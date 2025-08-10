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

func (driver *DB) FindTidyUrl(ctx context.Context) (*store.TidyUrl, error) {
	panic("not implemented")
}
