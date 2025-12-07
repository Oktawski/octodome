package qry

import (
	authdom "octodome.com/api/internal/auth/domain"
	"octodome.com/api/internal/core"
)

type EquipmentGetByID struct {
	ID   uint
	User authdom.UserContext
}

type EquipmentGetList struct {
	Pagination core.Pagination
	User       authdom.UserContext
}
