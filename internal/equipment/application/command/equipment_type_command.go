package eqcommand

import authdom "octodome/internal/auth/domain"

type CreateCommand struct {
	User authdom.UserContext
	Name string
}

type DeleteCommand struct {
	User authdom.UserContext
	ID   uint
}
