package events

import (
	"reflect"

	"octodome.com/shared/valuetype"
)

type MagicCodeRequested struct {
	Name  string
	Email valuetype.Email
	Code  string
}

func (e MagicCodeRequested) GetEventType() EventType {
	return EventType(reflect.TypeOf(e).Name())
}
