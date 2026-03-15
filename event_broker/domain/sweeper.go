package domain

import "context"

type Sweeper interface {
	Sweep(ctx context.Context) error
}
