package domain

type RoleName string

const (
	RoleUser  RoleName = "user"
	RoleAdmin RoleName = "admin"
)

var AvailableRolesStr = []string{
	string(RoleUser),
	string(RoleAdmin),
}

type RoleDTO struct {
	Name RoleName
}
