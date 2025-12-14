package domain

import (
	"context"

	"octodome.com/shared/events"
)

type Publisher interface {
	Publish(ctx context.Context, eventType events.EventType, payload interface{}) error
}
