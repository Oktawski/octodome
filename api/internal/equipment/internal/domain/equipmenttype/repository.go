package domain

import (
	authdom "octodome.com/api/internal/auth/domain"
)

type Repository interface {
	GetList(page int, pageSize int, user authdom.UserContext) ([]EquipmentType, int64, error)
	GetByID(id uint, user authdom.UserContext) (*EquipmentType, error)
	Create(eq *EquipmentType) error
	Update(eq *EquipmentType) error
	Delete(id uint) error

	ExistsByName(name string, user authdom.UserContext) bool
	IsUsed(id uint, user authdom.UserContext) bool
	OwnedByUser(id uint, user authdom.UserContext) bool
}
