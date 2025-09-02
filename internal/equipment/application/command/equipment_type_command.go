package eqcommand

import authdom "octodome/internal/auth/domain"

type EquipmentTypeCreateCommand struct {
	Name        string
	UserContext authdom.UserContext
}

type EquipmentTypeUpdateCommand struct {
	ID          uint
	Name        string
	UserContext authdom.UserContext
}

type EquipmentTypeDeleteCommand struct {
	ID          uint
	UserContext authdom.UserContext
}
