package domain

import "octodome.com/shared/valuetype"

type UserAuthDTO struct {
	ID       uint
	Email    valuetype.Email
	Password string
}
