package user

import (
	"octodome.com/api/internal/user/domain"
)

type GetByID struct {
	ID uint
}

type GetByIDHandler struct {
	repo domain.Repository
}

func NewUserGetByIDHandler(repo domain.Repository) *GetByIDHandler {
	return &GetByIDHandler{repo: repo}
}

func (handler *GetByIDHandler) Handle(q GetByID) (*domain.UserDTO, error) {
	user, err := handler.repo.GetByID(q.ID)
	if err != nil {
		return nil, err
	}

	return &domain.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, err
}
