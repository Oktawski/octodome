package events

import (
	"time"

	"octodome.com/shared/valuetype"
)

type UserRegistered struct {
	UserID       uint
	Email        valuetype.Email
	Name         string
	RegisteredAt time.Time
}
