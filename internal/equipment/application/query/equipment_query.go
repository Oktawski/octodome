package eqquery

import (
	authdom "octodome/internal/auth/domain"
	"octodome/internal/core"
)

type EquipmentGetByID struct {
	ID   uint
	User authdom.UserContext
}

type EquipmentGetList struct {
	Pagination core.Pagination
	User       authdom.UserContext
}
