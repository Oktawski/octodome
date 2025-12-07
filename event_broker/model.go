package eventbroker

import (
	"encoding/json"
	"time"
)

type EventStatus string

const (
	EventStatusPending   EventStatus = "pending"
	EventStatusProcessed EventStatus = "processed"
	EventStatusFailed    EventStatus = "failed"
)

type Event struct {
	ID        uint            `json:"id"`
	Type      string          `json:"type"`
	Payload   json.RawMessage `gorm:"type:jsonb"`
	CreatedAt time.Time       `json:"created_at"`
	Status    string          `json:"status"`
}

func (e *Event) TableName() string {
	return "events"
}
