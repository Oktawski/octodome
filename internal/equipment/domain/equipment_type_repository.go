package eqdom

import (
	authdom "octodome/internal/auth/domain"
)

type EquipmentTypeRepository interface {
	GetEquipmentTypes(page int, pageSize int, user authdom.UserContext) ([]EquipmentType, int64, error)
	GetEquipmentType(id uint, user authdom.UserContext) (*EquipmentType, error)
	CreateType(eq *EquipmentType) error
	DeleteEquipmentType(id uint, user authdom.UserContext) error
	ExistsByName(name string, user authdom.UserContext) bool
	IsUsed(id uint, user authdom.UserContext) bool
}
