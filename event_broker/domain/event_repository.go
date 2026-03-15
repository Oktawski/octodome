package domain

import "context"

type EventRepository interface {
	Save(ctx context.Context, eventType string, payload []byte) (Event, error)
	Get(ctx context.Context, eventType string) (Event, error)
	GetStale(ctx context.Context) ([]Event, error)
}
