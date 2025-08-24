package cqrs

import authdom "octodome/internal/auth/domain"

type WithUserContext struct {
	User authdom.UserContext
}
