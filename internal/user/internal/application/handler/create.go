package hdl

import (
	"octodome/internal/user/domain"
	cmd "octodome/internal/user/internal/application/command"

	"golang.org/x/crypto/bcrypt"
)

type CreateHandler struct {
	repo domain.Repository
}

func NewCreateHandler(repository domain.Repository) *CreateHandler {
	return &CreateHandler{repo: repository}
}

func (handler *CreateHandler) Handle(c cmd.Create) error {
	passwordHash, err := handler.hashPassword(c.Password)
	if err != nil {
		return err
	}

	userModel := &domain.User{
		Username:     c.Name,
		Email:        c.Email,
		PasswordHash: passwordHash,
	}

	return handler.repo.Create(userModel)
}

func (handler *CreateHandler) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
