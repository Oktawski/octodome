package equipmentdom

import authdom "octodome/internal/auth/domain"

type EquipmentRepository interface {
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
