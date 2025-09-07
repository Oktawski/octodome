package eqcmd

import (
	authdom "octodome/internal/auth/domain"
)

type Create struct {
	Name            string
	Description     string
	Category        string
	EquipmentTypeID uint
	UserContext     authdom.UserContext
}

type Update struct {
	ID          uint
	Name        string
	Description string
	Category    string
	UserContext authdom.UserContext
}

type Delete struct {
	ID          uint
	UserContext authdom.UserContext
}
