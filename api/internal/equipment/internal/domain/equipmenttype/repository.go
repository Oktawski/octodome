package domain

import (
	"context"

	authdom "octodome.com/api/internal/auth/domain"
)

type Repository interface {
	GetList(userContext authdom.UserContext, page int, pageSize int) ([]EquipmentType, int64, error)
	GetByID(userContext authdom.UserContext, id uint) (*EquipmentType, error)
	Create(eq *EquipmentType) error
	Update(userContext authdom.UserContext, ctx context.Context, eq *EquipmentType) error
	Delete(id uint) error

	ExistsByName(userContext authdom.UserContext, name string) bool
	IsUsed(userContext authdom.UserContext, id uint) bool
	OwnedByUser(userContext authdom.UserContext, id uint) bool
}
