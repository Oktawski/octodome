package cmd

import auth "octodome/internal/auth/domain"

type EquipmentCreate struct {
	Name            string
	Description     string
	Category        string
	EquipmentTypeID uint
	UserContext     auth.UserContext
}

type EquipmentUpdate struct {
	ID          uint
	Name        string
	Description string
	Category    string
	UserContext auth.UserContext
}

type EquipmentDelete struct {
	ID          uint
	UserContext auth.UserContext
}
