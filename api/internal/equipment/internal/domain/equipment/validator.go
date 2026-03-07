package domain

import authdom "octodome.com/api/internal/auth/domain"

type Validator interface {
	CanBeCreated(
		userContext authdom.UserContext,
		name string,
		equipmentTypeID uint,
	) bool

	CanBeModified(
		userContext authdom.UserContext,
		equipmentID uint,
	) bool
}

type validator struct {
	repo Repository
}

func NewValidator(repo Repository) *validator {
	return &validator{repo: repo}
}

func (v validator) CanBeCreated(
	userContext authdom.UserContext,
	name string,
	equipmentTypeID uint,
) bool {
	return !v.repo.ExistsByNameAndType(userContext, name, equipmentTypeID)
}

func (v validator) CanBeModified(
	userContext authdom.UserContext,
	equipmentID uint,
) bool {
	return v.repo.IsOwnedByUser(userContext, equipmentID)
}
