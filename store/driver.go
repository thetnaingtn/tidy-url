package store

import (
	"context"
)

type Driver interface {
	Create(ctx context.Context, p *TidyUrl) error
	FindTidyUrl(ctx context.Context, filters *Filters) (*TidyUrl, error)
	Close() error
}
