package model

import (
	"encoding/json"
	"errors"
	"time"

	"octodome.com/shared/events"
)

type Event struct {
	ID        uint            `json:"id"`
	Type      string          `json:"type"`
	Payload   json.RawMessage `gorm:"type:jsonb"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Status    string          `json:"status"`
}

func (e *Event) Pending() error {
	if e.Status == string(events.EventStatusFailed) ||
		e.Status == string(events.EventStatusProcessing) {
		e.Status = string(events.EventStatusPending)
		e.UpdatedAt = time.Now()
		return nil
	}
	return errors.New("cannot mark event as pending, it must be failed or processing")
}

func (e *Event) Processed() error {
	if e.Status == string(events.EventStatusProcessing) {
		e.Status = string(events.EventStatusProcessed)
		e.UpdatedAt = time.Now()
		return nil
	}
	return errors.New("cannot mark event as processed, it is not processing")
}

func (e *Event) Failed() error {
	if e.Status == string(events.EventStatusProcessing) {
		e.Status = string(events.EventStatusFailed)
		e.UpdatedAt = time.Now()
		return nil
	}
	return errors.New("cannot mark event as failed, it is not processing")
}

func (e *Event) Processing() error {
	if e.Status == string(events.EventStatusPending) ||
		e.Status == string(events.EventStatusFailed) {
		e.Status = string(events.EventStatusProcessing)
		e.UpdatedAt = time.Now()
		return nil
	}
	return errors.New("cannot mark event as processing, it must be pending or failed")
}

func (e *Event) TableName() string {
	return "events"
}
