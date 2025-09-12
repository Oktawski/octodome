package qry

import (
	authdom "octodome/internal/auth/domain"
	"octodome/internal/core"
)

type EquipmentTypeGetByID struct {
	ID   uint
	User authdom.UserContext
}

type EquipmentTypeGetList struct {
	Pagination core.Pagination
	User       authdom.UserContext
}
