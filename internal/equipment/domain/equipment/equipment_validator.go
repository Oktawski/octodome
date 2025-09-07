package eqdom

import authdom "octodome/internal/auth/domain"

type Validator interface {
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

type validator struct {
	repo Repository
}

func NewValidator(repo Repository) *validator {
	return &validator{repo: repo}
}

func (v validator) CanBeCreated(
	name string,
	equipmentTypeID uint,
	userContext authdom.UserContext,
) bool {
	return !v.repo.ExistsByNameAndType(name, equipmentTypeID, userContext)
}

func (v validator) CanBeModified(
	equipmentID uint,
	userContext authdom.UserContext,
) bool {
	return v.repo.IsOwnedByUser(equipmentID, userContext)
}
