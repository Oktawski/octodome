package repository

import (
	"context"
)

type MagicCode interface {
	Create(ctx context.Context, code string, email string) error
	GetByEmailAndCode(ctx context.Context, email string, code string) (string, error)
	DeleteByEmail(ctx context.Context, email string) error
}
