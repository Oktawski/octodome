package validator

import (
	"octodome/internal/auth/domain"
	"slices"
)

var validRoles = []domain.RoleName{
	domain.RoleUser,
	domain.RoleAdmin,
}

func CanBeUsed(role domain.RoleName) bool {
	return slices.Contains(validRoles, role)
}
