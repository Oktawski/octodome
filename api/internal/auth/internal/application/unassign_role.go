package auth

import (
	domainshared "octodome.com/api/internal/auth/domain"
	"octodome.com/api/internal/auth/internal/dependencies"
	domain "octodome.com/api/internal/auth/internal/domain/repository"
	"octodome.com/api/internal/auth/internal/domain/validator"
)

type UnassignRoleCommand struct {
	Role        domainshared.RoleName
	UserID      uint
	UserContext domainshared.UserContext
}

type UnassignRoleHandler interface {
	Handle(cmd UnassignRoleCommand) error
}

type unassignRoleHandler struct {
	repo      domain.RoleRepository
	validator validator.Role
}

func NewUnassignRoleHandler(deps dependencies.Container) UnassignRoleHandler {
	return &unassignRoleHandler{repo: deps.RoleRepository, validator: deps.RoleValidator}
}

func (h *unassignRoleHandler) Handle(cmd UnassignRoleCommand) error {
	if err := h.validator.CanBeUnassigned(
		cmd.Role,
		cmd.UserID,
		cmd.UserContext); err != nil {
		return err
	}

	err := h.repo.UnassignRole(cmd.Role, cmd.UserID)

	return err
}
