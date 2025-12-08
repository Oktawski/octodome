package events

type EventType string

type EventStatus string

const (
	EventStatusPending    EventStatus = "pending"
	EventStatusProcessing EventStatus = "processing"
	EventStatusProcessed  EventStatus = "processed"
	EventStatusFailed     EventStatus = "failed"
)
