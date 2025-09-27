package auth

import (
	"fmt"
	domainshared "octodome/internal/auth/domain"
	"octodome/internal/auth/internal/dependencies"
	domain "octodome/internal/auth/internal/domain/repository"
	"octodome/internal/auth/internal/domain/validator"
	"strings"
)

type AssignRoleCommand struct {
	Role        domainshared.RoleName
	UserID      uint
	UserContext domainshared.UserContext
}

type AssignRoleHandler interface {
	Handle(cmd AssignRoleCommand) error
}

type assignRoleHandler struct {
	repo      domain.RoleRepository
	validator validator.Role
}

func NewAssignRoleHandler(deps dependencies.Container) AssignRoleHandler {
	return &assignRoleHandler{
		repo:      deps.RoleRepository,
		validator: deps.RoleValidator,
	}
}

func (h *assignRoleHandler) Handle(c AssignRoleCommand) error {
	if !h.validator.CanBeUsed(c.Role) {
		return fmt.Errorf(
			"role cannot be used. available roles: %s",
			strings.Join(domainshared.AvailableRolesStr, ", "))
	}

	err := h.repo.AssignRole(c.Role, c.UserID)
	return err
}
