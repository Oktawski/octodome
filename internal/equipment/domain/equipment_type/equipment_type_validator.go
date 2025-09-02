package eqtypedom

import authdom "octodome/internal/auth/domain"

type EquipmentTypeValidator interface {
	CanBeCreated(name string, userContext authdom.UserContext) bool
	CanBeModified(id uint, userContext authdom.UserContext) bool
}

type equipmentTypeValidator struct {
	repo EquipmentTypeRepository
}

func NewEquipmentTypeValidator(repo EquipmentTypeRepository) *equipmentTypeValidator {
	return &equipmentTypeValidator{repo: repo}
}

func (v equipmentTypeValidator) CanBeCreated(
	name string,
	userContext authdom.UserContext) bool {

	return !v.repo.ExistsByName(name, userContext)
}

func (v equipmentTypeValidator) CanBeModified(
	id uint,
	userContext authdom.UserContext,
) bool {
	return !v.repo.IsUsed(id, userContext) &&
		v.repo.OwnedByUser(id, userContext)
}
