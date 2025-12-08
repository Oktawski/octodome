package qry

import (
	authdom "octodome.com/api/internal/auth/domain"
	"octodome.com/shared/collection"
)

type EquipmentTypeGetByID struct {
	ID   uint
	User authdom.UserContext
}

type EquipmentTypeGetList struct {
	Pagination collection.Pagination
	User       authdom.UserContext
}
