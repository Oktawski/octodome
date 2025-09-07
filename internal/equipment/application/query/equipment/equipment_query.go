package eqqry

import (
	authdom "octodome/internal/auth/domain"
	"octodome/internal/core"
)

type GetByID struct {
	ID   uint
	User authdom.UserContext
}

type GetList struct {
	Pagination core.Pagination
	User       authdom.UserContext
}
