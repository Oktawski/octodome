package hdl

import (
	qry "octodome/internal/user/application/query"
	domain "octodome/internal/user/domain"
)

type GetByIDHandler struct {
	repo domain.Repository
}

func NewUserGetByIDHandler(repo domain.Repository) *GetByIDHandler {
	return &GetByIDHandler{repo: repo}
}

func (handler *GetByIDHandler) Handle(q qry.GetByID) (*domain.UserDTO, error) {
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
