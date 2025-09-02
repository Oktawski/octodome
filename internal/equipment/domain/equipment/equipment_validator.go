package equipmentdom

import authdom "octodome/internal/auth/domain"

type EquipmentValidator interface {
	CanBeCreated(
		name string,
		equipmentTypeID uint,
		userContext authdom.UserContext,
	) bool

	CanBeModified(
		equipmentID uint,
		userContext authdom.UserContext,
	) bool
}

type equipmentValidator struct {
	repo EquipmentRepository
}

func NewEquipmentValidator(repo EquipmentRepository) *equipmentValidator {
	return &equipmentValidator{repo: repo}
}

func (v equipmentValidator) CanBeCreated(
	name string,
	equipmentTypeID uint,
	userContext authdom.UserContext,
) bool {
	return !v.repo.ExistsByNameAndType(name, equipmentTypeID, userContext)
}

func (v equipmentValidator) CanBeModified(
	equipmentID uint,
	userContext authdom.UserContext,
) bool {
	return v.repo.IsOwnedByUser(equipmentID, userContext)
}
