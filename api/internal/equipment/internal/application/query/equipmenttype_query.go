package qry

import (
	authdom "octodome.com/api/internal/auth/domain"
	"octodome.com/api/internal/core"
)

type EquipmentTypeGetByID struct {
	ID   uint
	User authdom.UserContext
}

type EquipmentTypeGetList struct {
	Pagination core.Pagination
	User       authdom.UserContext
}
