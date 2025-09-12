package auth

import (
	"octodome/internal/auth/internal/domain"
)

type AuthenticateCommand struct {
	Username string
	Password string
}

type AuthenticateHandler interface {
	Handle(authReq *AuthenticateCommand) (string, error)
}

type authenticateHandler struct {
	userRepository domain.Repository
	tokenGenerator domain.AuthTokenGenerator
	passwordHasher domain.PasswordHasher
}

func NewAuthenticateHandler(
	repository domain.Repository,
	tokenGenerator domain.AuthTokenGenerator,
	passwordHasher domain.PasswordHasher) AuthenticateHandler {
	return &authenticateHandler{
		userRepository: repository,
		tokenGenerator: tokenGenerator,
		passwordHasher: passwordHasher,
	}
}

func (handler *authenticateHandler) Handle(request *AuthenticateCommand) (string, error) {
	user, err := handler.userRepository.GetUserByUsername(request.Username)
	if err != nil {
		return "", err
	}

	if err := handler.passwordHasher.CompareHashAndPassword(
		user.PasswordHash,
		request.Password); err != nil {
		return "", err
	}

	return handler.tokenGenerator.GenerateToken(user)
}
