package userhandler

import (
	userquery "octodome/internal/user/application/query"
	userdom "octodome/internal/user/domain"
)

type GetByIDHandler struct {
	userRepository userdom.UserRepository
}

func NewUserGetByIDHandler(repository userdom.UserRepository) *GetByIDHandler {
	return &GetByIDHandler{userRepository: repository}
}

func (handler *GetByIDHandler) Handle(q userquery.GetByID) (*userdom.UserDTO, error) {
	user, err := handler.userRepository.GetByID(q.ID)
	if err != nil {
		return nil, err
	}

	return &userdom.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, err
}
