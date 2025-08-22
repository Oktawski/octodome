package authdom

import userdom "octodome/internal/user/domain"

type AuthTokenGenerator interface {
	GenerateToken(user *userdom.User) (string, error)
}
