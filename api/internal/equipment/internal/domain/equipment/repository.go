package domain

import authdom "octodome.com/api/internal/auth/domain"

type Repository interface {
	GetList(page, pageSize int, user authdom.UserContext) ([]Equipment, int64, error)
	GetByID(id uint, user authdom.UserContext) (*Equipment, error)
	Create(equipment *Equipment) error
	Update(equipment *Equipment) error
	Delete(id uint) error

	ExistsByNameAndType(
		name string,
		equipmentTypeID uint,
		userContext authdom.UserContext) bool

	IsOwnedByUser(equipmentID uint, userContext authdom.UserContext) bool
}
