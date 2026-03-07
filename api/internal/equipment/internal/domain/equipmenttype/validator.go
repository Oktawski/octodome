package domain

import authdom "octodome.com/api/internal/auth/domain"

type Validator interface {
	CanBeCreated(userContext authdom.UserContext, name string) bool
	CanBeModified(userContext authdom.UserContext, id uint) bool
}

type validator struct {
	repo Repository
}

func NewEquipmentTypeValidator(repo Repository) *validator {
	return &validator{repo: repo}
}

func (v validator) CanBeCreated(
	userContext authdom.UserContext,
	name string,
) bool {

	return !v.repo.ExistsByName(userContext, name)
}

func (v validator) CanBeModified(
	userContext authdom.UserContext,
	id uint,
) bool {
	if v.repo.IsUsed(userContext, id) {
		return false
	}

	return userContext.HasRole(authdom.RoleAdmin) ||
		v.repo.OwnedByUser(userContext, id)
}
