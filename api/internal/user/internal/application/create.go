package user

import (
	"context"

	"octodome.com/api/internal/user/domain"

	"golang.org/x/crypto/bcrypt"
)

type Create struct {
	Context  context.Context
	Name     string
	Email    string
	Password string
}

type CreateHandler struct {
	repo domain.Repository
}

func NewCreateHandler(
	repository domain.Repository,
) *CreateHandler {
	return &CreateHandler{repo: repository}
}

func (handler *CreateHandler) Handle(c Create) error {
	passwordHash, err := handler.hashPassword(c.Password)
	if err != nil {
		return err
	}

	userModel := &domain.User{
		Username:     c.Name,
		Email:        c.Email,
		PasswordHash: passwordHash,
	}

	_, err = handler.repo.Create(c.Context, userModel)
	if err != nil {
		return err
	}

	return nil
}

func (handler *CreateHandler) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
