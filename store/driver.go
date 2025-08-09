package store

import (
	"context"
)

type Driver interface {
	Create(ctx context.Context, p *TidyUrl) error
	FindTidyUrl(ctx context.Context) (*TidyUrl, error)
}
