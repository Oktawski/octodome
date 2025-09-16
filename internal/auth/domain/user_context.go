package domain

import (
	"octodome/internal/core/collection"
)

type ContextKey string

const UserContextKey ContextKey = "user"

type UserContext struct {
	ID    uint
	Roles []RoleDTO
}

func (uc *UserContext) HasRole(roleName RoleName) bool {
	for _, role := range uc.Roles {
		if role.Name == roleName {
			return true
		}
	}
	return false
}

func (uc *UserContext) HasRoleAny(roles []RoleName) bool {
	if uc == nil || len(roles) == 0 || len(uc.Roles) == 0 {
		return false
	}

	allowed := collection.ToSet(roles)

	for _, userRole := range uc.Roles {
		if _, ok := allowed[userRole.Name]; ok {
			return true
		}
	}

	return false
}

var roleRank = map[RoleName]int{
	RoleUser:  1,
	RoleAdmin: 42,
}

func (uc *UserContext) HasAtLeastRole(required RoleName) bool {
	reqRank, ok := roleRank[required]
	if !ok || uc == nil {
		return false
	}
	for _, r := range uc.Roles {
		if roleRank[r.Name] >= reqRank {
			return true
		}
	}
	return false
}
