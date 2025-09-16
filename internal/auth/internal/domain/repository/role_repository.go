package domain

import "octodome/internal/auth/domain"

type RoleRepository interface {
	GetRolesByUserID(userID uint) ([]domain.RoleDTO, error)
}
