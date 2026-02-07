package events

import (
	"reflect"
	"time"

	"octodome.com/shared/valuetype"
)

type UserRegistered struct {
	UserID       uint
	Email        valuetype.Email
	Name         string
	RegisteredAt time.Time
}

func (e UserRegistered) GetEventType() EventType {
	return EventType(reflect.TypeOf(e).Name())
}
