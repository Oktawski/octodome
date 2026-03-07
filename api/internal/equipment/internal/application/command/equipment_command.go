package cmd

import (
	"context"

	auth "octodome.com/api/internal/auth/domain"
)

type EquipmentCreate struct {
	UserContext     auth.UserContext
	Name            string
	Description     string
	Category        string
	EquipmentTypeID uint
}

type EquipmentUpdate struct {
	UserContext auth.UserContext
	Ctx         context.Context
	ID          uint
	Name        string
	Description string
	Category    string
}

type EquipmentDelete struct {
	ID          uint
	UserContext auth.UserContext
}
