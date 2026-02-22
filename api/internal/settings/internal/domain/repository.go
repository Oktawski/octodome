package domain

import "context"

type Repository interface {
	Get(ctx context.Context, name string) (*Setting, error)
	Set(ctx context.Context, setting Setting) error
}
