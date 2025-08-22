package eqdom

import (
	authdom "octodome/internal/auth/domain"
)

type EquipmentRepository interface {
	GetEquipmentTypes(user *authdom.UserContext) (*[]EquipmentType, error)
	GetEquipmentType(id uint, user *authdom.UserContext) (*EquipmentType, error)
	CreateType(eq *EquipmentType) error
}
