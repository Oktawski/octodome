package events

import "time"

type UserRegistered struct {
	UserID       uint
	Email        string
	Name         string
	RegisteredAt time.Time
}
