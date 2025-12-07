package cmd

import authdom "octodome.com/api/internal/auth/domain"

type EquipmentTypeCreate struct {
	Name        string
	UserContext authdom.UserContext
}

type EquipmentTypeUpdate struct {
	ID          uint
	Name        string
	UserContext authdom.UserContext
}

type EquipmentTypeDelete struct {
	ID          uint
	UserContext authdom.UserContext
}
