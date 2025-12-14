package auth

import (
	"context"
	"sync"

	domainshared "octodome.com/api/internal/auth/domain"
	"octodome.com/api/internal/auth/internal/dependencies"
	domain "octodome.com/api/internal/auth/internal/domain/repository"
	"octodome.com/api/internal/auth/internal/domain/validator"
	"octodome.com/shared/collection"
)

type SyncRolesCommand struct {
	Context     context.Context
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

	currentRoleDTOs, err := h.repo.GetRolesByUserID(cmd.Context, cmd.UserID)
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

	return h.repo.SyncRoles(cmd.Context, rolesToAdd, rolesToRemove, cmd.UserID)
}

func getRolesToAddAndDelete(
	current, desired map[domainshared.RoleName]struct{},
) ([]domainshared.RoleName, []domainshared.RoleName) {
	var toAdd, toRemove []domainshared.RoleName

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for role := range desired {
			if _, found := current[role]; !found {
				toAdd = append(toAdd, role)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for role := range current {
			if _, found := desired[role]; !found {
				toRemove = append(toRemove, role)
			}
		}
	}()

	wg.Wait()
	return toAdd, toRemove
}
