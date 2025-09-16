package domain

import authdom "octodome/internal/auth/domain"

type Validator interface {
	CanBeCreated(name string, userContext authdom.UserContext) bool
	CanBeModified(id uint, userContext authdom.UserContext) bool
}

type validator struct {
	repo Repository
}

func NewEquipmentTypeValidator(repo Repository) *validator {
	return &validator{repo: repo}
}

func (v validator) CanBeCreated(
	name string,
	userContext authdom.UserContext) bool {

	return !v.repo.ExistsByName(name, userContext)
}

func (v validator) CanBeModified(
	id uint,
	userContext authdom.UserContext,
) bool {
	if v.repo.IsUsed(id, userContext) {
		return false
	}

	return userContext.HasRole(authdom.RoleAdmin) ||
		v.repo.OwnedByUser(id, userContext)
}
