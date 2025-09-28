package auth

import (
	domainshared "octodome/internal/auth/domain"
	"octodome/internal/auth/internal/dependencies"
	domain "octodome/internal/auth/internal/domain/repository"
	"octodome/internal/auth/internal/domain/validator"
	"octodome/internal/core/collection"
)

type SyncRolesCommand struct {
	Roles       []domainshared.RoleName
	UserID      uint
	UserContext domainshared.UserContext
}

type SyncRolesHandler interface {
	Handle(cmd SyncRolesCommand) error
}

type syncRolesHandler struct {
	repo      domain.RoleRepository
	validator validator.Role
}

func NewSyncRolesHandler(deps dependencies.Container) *syncRolesHandler {
	return &syncRolesHandler{
		repo:      deps.RoleRepository,
		validator: deps.RoleValidator,
	}
}

func (h *syncRolesHandler) Handle(cmd SyncRolesCommand) error {
	for _, role := range cmd.Roles {
		if err := h.validator.CanBeUsed(role); err != nil {
			return err
		}
	}

	currentRoleDTOs, err := h.repo.GetRolesByUserID(cmd.UserID)
	if err != nil {
		return err
	}

	currentRoles := collection.Map(
		currentRoleDTOs,
		func(e domainshared.RoleDTO) domainshared.RoleName {
			return e.Name
		},
	)
	currentRolesSet := collection.ToSet(currentRoles)
	desiredRolesMap := collection.ToSet(cmd.Roles)

	rolesToAdd, rolesToRemove := getRolesToAddAndDelete(
		currentRolesSet,
		desiredRolesMap,
	)

	return h.repo.SyncRoles(rolesToAdd, rolesToRemove, cmd.UserID)
}

func getRolesToAddAndDelete(
	current, desired map[domainshared.RoleName]struct{},
) ([]domainshared.RoleName, []domainshared.RoleName) {
	var toAdd, toRemove []domainshared.RoleName
	for role := range desired {
		if _, found := current[role]; !found {
			toAdd = append(toAdd, role)
		}
	}

	for role := range current {
		if _, found := desired[role]; !found {
			toRemove = append(toRemove, role)
		}
	}

	return toAdd, toRemove
}
