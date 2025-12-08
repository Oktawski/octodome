package domain

import "octodome.com/shared/events"

type Publisher interface {
	Publish(eventType events.EventType, payload interface{}) error
}
