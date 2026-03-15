package domain

import "context"

type Dispatcher interface {
	Dispatch(ctx context.Context, event Event, handler Handler) error
}
