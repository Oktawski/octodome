package validator

import (
	"errors"
	"fmt"
	domainshared "octodome/internal/auth/domain"
	domain "octodome/internal/auth/internal/domain/repository"
	"slices"
	"strings"
)

var validRoles = []domainshared.RoleName{
	domainshared.RoleUser,
	domainshared.RoleAdmin,
}

type Role interface {
	CanBeUsed(roleName domainshared.RoleName) error
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

func (r *role) CanBeUsed(role domainshared.RoleName) error {
	if !slices.Contains(validRoles, role) {
		return fmt.Errorf(
			"role cannot be used. available roles: %s",
			strings.Join(domainshared.AvailableRolesStr, ", "))
	}
	return nil
}

func (r *role) CanBeUnassigned(
	roleName domainshared.RoleName,
	userID uint,
	userContext domainshared.UserContext,
) error {
	if userID == userContext.ID {
		return errors.New("cannot unassign role for self")
	}
	if !userContext.HasRole(domainshared.RoleAdmin) {
		return errors.New("only admin can unassign roles")
	}
	return nil
}
