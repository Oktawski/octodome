package auth

import (
	authdom "octodome/internal/auth/domain"
	userdom "octodome/internal/user/domain"

	"golang.org/x/crypto/bcrypt"
)

type AuthenticateHandler interface {
	Handle(authReq *AuthenticateCommand) (string, error)
}

type AuthRepository interface {
	GetUserByUsername(username string) (user *userdom.User, err error)
}

type authenticateHandler struct {
	userRepository AuthRepository
	tokenGenerator authdom.AuthTokenGenerator
}

func NewAuthenticateHandler(
	repository AuthRepository,
	tokenGenerator authdom.AuthTokenGenerator) AuthenticateHandler {
	return &authenticateHandler{
		userRepository: repository,
		tokenGenerator: tokenGenerator,
	}
}

func (handler *authenticateHandler) Handle(request *AuthenticateCommand) (string, error) {
	user, err := handler.userRepository.GetUserByUsername(request.Username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(request.Password)); err != nil {
		return "", err
	}

	return handler.tokenGenerator.GenerateToken(user)
}
