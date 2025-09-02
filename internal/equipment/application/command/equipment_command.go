package eqcommand

import (
	authdom "octodome/internal/auth/domain"
)

type EquipmentCreateCommand struct {
	Name            string
	Description     string
	Category        string
	EquipmentTypeID uint
	UserContext     authdom.UserContext
}

type EquipmentUpdateCommand struct {
	ID          uint
	Name        string
	Description string
	Category    string
	UserContext authdom.UserContext
}

type EquipmentDeleteCommand struct {
	ID          uint
	UserContext authdom.UserContext
}
