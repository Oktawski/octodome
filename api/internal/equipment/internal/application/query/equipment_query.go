package qry

import (
	authdom "octodome.com/api/internal/auth/domain"
	"octodome.com/shared/collection"
)

type EquipmentGetByID struct {
	ID   uint
	User authdom.UserContext
}

type EquipmentGetList struct {
	Pagination collection.Pagination
	User       authdom.UserContext
}
