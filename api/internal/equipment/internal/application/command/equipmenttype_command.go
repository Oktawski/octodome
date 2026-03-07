package cmd

import (
	"context"

	authdom "octodome.com/api/internal/auth/domain"
)

type EquipmentTypeCreate struct {
	Name        string
	UserContext authdom.UserContext
}

type EquipmentTypeUpdate struct {
	UserContext authdom.UserContext
	Ctx         context.Context
	ID          uint
	Name        string
}

type EquipmentTypeDelete struct {
	ID          uint
	UserContext authdom.UserContext
}
