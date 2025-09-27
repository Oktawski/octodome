package dependencies

import (
	"octodome/internal/auth/internal/domain"
	"octodome/internal/auth/internal/domain/repository"
	"octodome/internal/auth/internal/domain/validator"
)

type Container struct {
	UserReader     repository.UserReader
	RoleRepository repository.RoleRepository
	TokenGenerator domain.AuthTokenGenerator
	PasswordHasher domain.PasswordHasher
	RoleValidator  validator.Role
}

func NewContainer(
	userReader repository.UserReader,
	roleRepository repository.RoleRepository,
	tokenGenerator domain.AuthTokenGenerator,
	passwordHasher domain.PasswordHasher,
	roleValidator validator.Role,
) Container {
	return Container{
		UserReader:     userReader,
		RoleRepository: roleRepository,
		TokenGenerator: tokenGenerator,
		PasswordHasher: passwordHasher,
		RoleValidator:  roleValidator,
	}
}
