package userhandler

import (
	usercommand "octodome/internal/user/application/command"
	userdom "octodome/internal/user/domain"

	"golang.org/x/crypto/bcrypt"
)

type CreateHandler struct {
	userRepository userdom.UserRepository
}

func NewUserCreateHandler(repository userdom.UserRepository) *CreateHandler {
	return &CreateHandler{userRepository: repository}
}

func (handler *CreateHandler) Handle(c usercommand.Create) error {
	passwordHash, err := handler.hashPassword(c.Password)
	if err != nil {
		return err
	}

	userModel := &userdom.User{
		Username:     c.Name,
		Email:        c.Email,
		PasswordHash: passwordHash,
	}

	return handler.userRepository.Create(userModel)
}

func (handler *CreateHandler) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
