package domain

import (
	"context"

	authdom "octodome.com/api/internal/auth/domain"
)

type Repository interface {
	GetList(userContext authdom.UserContext, page int, pageSize int) ([]Equipment, int64, error)
	GetByID(userContext authdom.UserContext, id uint) (*Equipment, error)
	Create(equipment *Equipment) error

	Update(userContext authdom.UserContext,
		ctx context.Context,
		equipment *Equipment) error

	Delete(id uint) error
	ExistsByNameAndType(
		userContext authdom.UserContext,
		name string,
		equipmentTypeID uint) bool
	IsOwnedByUser(userContext authdom.UserContext, equipmentID uint) bool
}
