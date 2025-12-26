package infra

import (
	"context"
	"encoding/json"
	"time"

	"gorm.io/gorm"
	"octodome.com/eventbroker/domain"
	"octodome.com/shared/events"
)

type publisher struct {
	db *gorm.DB
}

func NewEventPublisher(db *gorm.DB) domain.Publisher {
	return &publisher{db: db}
}

func (p *publisher) Publish(
	ctx context.Context,
	eventType string,
	payload interface{},
) error {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	eventModel := &event{
		Type:      eventType,
		Payload:   payloadJSON,
		CreatedAt: time.Now(),
		Status:    string(events.EventStatusPending),
	}

	return gorm.G[event](p.db).Create(ctx, eventModel)
}
