package eqquery

import (
	authdom "octodome/internal/auth/domain"
)

type GetListQuery struct {
	User     authdom.UserContext
	Page     int
	PageSize int
}

type GetQuery struct {
	User authdom.UserContext
	ID   uint
}
