package publisher

import (
	"encoding/json"
	"reflect"
	"time"

	"gorm.io/gorm"
	"octodome.com/eventbroker"
)

type Publisher interface {
	Publish(event interface{}) error
}

type publisher struct {
	db *gorm.DB
}

func NewEventPublisher(db *gorm.DB) Publisher {
	return &publisher{db: db}
}

func (p *publisher) Publish(payload interface{}) error {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	event := &eventbroker.Event{
		Type:      reflect.TypeOf(payload).Name(),
		Payload:   payloadJSON,
		CreatedAt: time.Now(),
		Status:    string(eventbroker.EventStatusPending),
	}

	return p.db.Create(&event).Error
}
