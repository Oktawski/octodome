package infra

import (
	"encoding/json"
	"errors"
	"time"

	"octodome.com/shared/events"
)

type event struct {
	ID        uint            `json:"id"`
	Type      string          `json:"type"`
	Payload   json.RawMessage `gorm:"type:jsonb"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Status    string          `json:"status"`
}

func (e *event) Pending() error {
	if e.Status == string(events.EventStatusFailed) ||
		e.Status == string(events.EventStatusProcessing) {
		e.Status = string(events.EventStatusPending)
		e.UpdatedAt = time.Now()
		return nil
	}
	return errors.New("cannot mark event as pending, it must be failed or processing")
}

func (e *event) Processed() error {
	if e.Status == string(events.EventStatusProcessing) {
		e.Status = string(events.EventStatusProcessed)
		e.UpdatedAt = time.Now()
		return nil
	}
	return errors.New("cannot mark event as processed, it is not processing")
}

func (e *event) Failed() error {
	if e.Status == string(events.EventStatusProcessing) {
		e.Status = string(events.EventStatusFailed)
		e.UpdatedAt = time.Now()
		return nil
	}
	return errors.New("cannot mark event as failed, it is not processing")
}

func (e *event) Processing() error {
	if e.Status == string(events.EventStatusPending) ||
		e.Status == string(events.EventStatusFailed) {
		e.Status = string(events.EventStatusProcessing)
		e.UpdatedAt = time.Now()
		return nil
	}
	return errors.New("cannot mark event as processing, it must be pending or failed")
}

func (e *event) TableName() string {
	return "events"
}
