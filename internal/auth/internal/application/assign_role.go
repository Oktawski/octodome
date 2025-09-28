package auth

import (
	domainshared "octodome/internal/auth/domain"
	"octodome/internal/auth/internal/dependencies"
	domain "octodome/internal/auth/internal/domain/repository"
	"octodome/internal/auth/internal/domain/validator"
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
	if err := h.validator.CanBeUsed(c.Role); err != nil {
		return err
	}

	err := h.repo.AssignRole(c.Role, c.UserID)
	return err
}
