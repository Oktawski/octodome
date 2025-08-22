package auth

import (
	authdom "octodome/internal/auth/domain"
	userdom "octodome/internal/user/domain"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler interface {
	Authenticate(authReq *AuthenticateRequest) (string, error)
}

type AuthRepository interface {
	GetUserByUsername(username string) (user *userdom.User, err error)
}

type authHandler struct {
	userRepository AuthRepository
	tokenGenerator authdom.AuthTokenGenerator
}

func NewAuthHandler(
	repository AuthRepository,
	tokenGenerator authdom.AuthTokenGenerator) AuthHandler {
	return &authHandler{
		userRepository: repository,
		tokenGenerator: tokenGenerator,
	}
}

func (handler *authHandler) Authenticate(request *AuthenticateRequest) (string, error) {
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
