package auth

import (
	"octodome/internal/auth/internal/dependencies"
	"octodome/internal/auth/internal/domain"
	domainrepo "octodome/internal/auth/internal/domain/repository"
)

type AuthenticateCommand struct {
	Username string
	Password string
}

type AuthenticateHandler interface {
	Handle(authReq AuthenticateCommand) (string, error)
}

type authenticateHandler struct {
	userReader     domainrepo.UserReader
	tokenGenerator domain.AuthTokenGenerator
	passwordHasher domain.PasswordHasher
}

func NewAuthenticateHandler(deps dependencies.Container) AuthenticateHandler {
	return &authenticateHandler{
		userReader:     deps.UserReader,
		tokenGenerator: deps.TokenGenerator,
		passwordHasher: deps.PasswordHasher,
	}
}

func (handler *authenticateHandler) Handle(request AuthenticateCommand) (string, error) {
	userAuthDTO, err := handler.userReader.GetUserAuthDTO(request.Username)
	if err != nil {
		return "", err
	}

	if err := handler.passwordHasher.CompareHashAndPassword(
		userAuthDTO.Password,
		request.Password); err != nil {
		return "", err
	}

	return handler.tokenGenerator.GenerateToken(userAuthDTO)
}
