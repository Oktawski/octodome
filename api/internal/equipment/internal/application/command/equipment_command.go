package cmd

import (
	"context"

	auth "octodome.com/api/internal/auth/domain"
)

type EquipmentCreate struct {
	Ctx             context.Context
	UserContext     auth.UserContext
	Name            string
	Description     string
	Category        string
	EquipmentTypeID uint
}

type EquipmentUpdate struct {
	Ctx         context.Context
	UserContext auth.UserContext
	ID          uint
	Name        string
	Description string
	Category    string
}

type EquipmentDelete struct {
	Ctx         context.Context
	UserContext auth.UserContext
	ID          uint
}
