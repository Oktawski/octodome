package qry

import (
	authdom "octodome.com/api/internal/auth/domain"
	"octodome.com/shared/collection"
)

type EquipmentTypeGetByID struct {
	User authdom.UserContext
	ID   uint
}

type EquipmentTypeGetList struct {
	User       authdom.UserContext
	Pagination collection.Pagination
}
