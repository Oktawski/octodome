package user

import (
	userdom "octodome/internal/user/domain"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	GetUserByID(id uint) (*UserGetResponse, error)
	CreateUser(user *UserCreateRequest) error
}

type userHandler struct {
	userRepository userdom.UserRepository
}

func NewUserHandler(repository userdom.UserRepository) UserHandler {
	return &userHandler{userRepository: repository}
}

func (handler *userHandler) GetUserByID(id uint) (*UserGetResponse, error) {
	user, err := handler.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &UserGetResponse{
		ID:    user.ID,
		Name:  user.Username,
		Email: user.Email,
	}, err
}

func (handler *userHandler) CreateUser(user *UserCreateRequest) error {
	passwordHash, err := handler.hashPassword(user.Password)
	if err != nil {
		return err
	}

	userModel := &userdom.User{
		Username:     user.Name,
		Email:        user.Email,
		PasswordHash: passwordHash,
	}

	return handler.userRepository.Create(userModel)
}

func (handler *userHandler) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
