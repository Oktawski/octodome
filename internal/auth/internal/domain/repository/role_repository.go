package repository

import "octodome/internal/auth/domain"

type RoleRepository interface {
	GetRolesByUserID(userID uint) ([]domain.RoleDTO, error)
	AssignRole(role domain.RoleName, userID uint) error
	UnassignRole(role domain.RoleName, userID uint) error
	SyncRoles(rolesToAdd []domain.RoleName, rolesToRemove []domain.RoleName, userID uint) error
}
