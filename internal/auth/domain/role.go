package domain

type RoleName string

const (
	RoleUser  RoleName = "user"
	RoleAdmin RoleName = "admin"
)

type RoleDTO struct {
	Name RoleName
}
