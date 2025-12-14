package repository

import (
	"context"

	"octodome.com/api/internal/auth/domain"
)

type RoleRepository interface {
	GetRolesByUserID(ctx context.Context, userID uint) ([]domain.RoleDTO, error)
	AssignRole(ctx context.Context, role domain.RoleName, userID uint) error
	UnassignRole(ctx context.Context, role domain.RoleName, userID uint) error
	SyncRoles(ctx context.Context, rolesToAdd []domain.RoleName, rolesToRemove []domain.RoleName, userID uint) error
}
