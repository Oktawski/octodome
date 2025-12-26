package events

import "reflect"

type EventType any

type EventStatus string

const (
	EventStatusPending    EventStatus = "pending"
	EventStatusProcessing EventStatus = "processing"
	EventStatusProcessed  EventStatus = "processed"
	EventStatusFailed     EventStatus = "failed"
)

func GetEventTypeName(event EventType) EventType {
	return EventType(reflect.TypeOf(event).Name())
}
