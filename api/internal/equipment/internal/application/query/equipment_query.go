package qry

import (
	"context"

	authdom "octodome.com/api/internal/auth/domain"
	"octodome.com/shared/collection"
)

type EquipmentGetByID struct {
	Ctx         context.Context
	UserContext authdom.UserContext
	ID          uint
}

type EquipmentGetList struct {
	Ctx         context.Context
	UserContext authdom.UserContext
	Pagination  collection.Pagination
}
