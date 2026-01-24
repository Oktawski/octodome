package domain

import (
	"octodome.com/shared/valuetype"
)

type User struct {
	ID           uint
	Username     string
	Email        valuetype.Email
	PasswordHash string
}

type UserDTO struct {
	ID       uint
	Username string
	Email    valuetype.Email
}
