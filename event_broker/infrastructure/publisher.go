package infra

import (
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

func (p *publisher) Publish(eventType events.EventType, payload interface{}) error {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	event := &event{
		Type:      string(eventType),
		Payload:   payloadJSON,
		CreatedAt: time.Now(),
		Status:    string(events.EventStatusPending),
	}

	return p.db.Create(&event).Error
}
