package validator

import (
	"errors"
	domainshared "octodome/internal/auth/domain"
	domain "octodome/internal/auth/internal/domain/repository"
	"slices"
)

var validRoles = []domainshared.RoleName{
	domainshared.RoleUser,
	domainshared.RoleAdmin,
}

type Role interface {
	CanBeUsed(roleName domainshared.RoleName) bool
	CanBeUnassigned(
		roleName domainshared.RoleName,
		userID uint,
		userContext domainshared.UserContext) error
}

type role struct {
	repo domain.RoleRepository
}

func NewRoleValidator(repo domain.RoleRepository) Role {
	return &role{repo: repo}
}

func (r *role) CanBeUsed(role domainshared.RoleName) bool {
	return slices.Contains(validRoles, role)
}

func (r *role) CanBeUnassigned(
	roleName domainshared.RoleName,
	userID uint,
	userContext domainshared.UserContext,
) error {
	if userID == userContext.ID {
		return errors.New("cannot unassign role for self")
	}
	return nil
}
