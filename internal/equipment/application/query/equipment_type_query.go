package eqquery

import (
	authdom "octodome/internal/auth/domain"
)

type GetByID struct {
	User authdom.UserContext
	ID   uint
}

type GetList struct {
	User     authdom.UserContext
	Page     int
	PageSize int
}
