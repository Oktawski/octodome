package auth

import (
	"fmt"
	domainshared "octodome/internal/auth/domain"
	domain "octodome/internal/auth/internal/domain/repository"
	"octodome/internal/auth/internal/domain/validator"
	"strings"
)

type AssignRoleCommand struct {
	Role   domainshared.RoleName
	UserID uint
}

type AssignRoleHandler interface {
	Handle(cmd AssignRoleCommand) error
}

type assignRoleHandler struct {
	repo domain.RoleRepository
}

func NewAssignRoleHandler(
	repo domain.RoleRepository,
) AssignRoleHandler {
	return &assignRoleHandler{repo: repo}
}

func (h *assignRoleHandler) Handle(c AssignRoleCommand) error {
	if !validator.CanBeUsed(c.Role) {
		return fmt.Errorf(
			"role cannot be used. available roles: %s",
			strings.Join(domainshared.AvailableRolesStr, ", "))
	}

	err := h.repo.AssignRole(c.Role, c.UserID)
	return err
}
