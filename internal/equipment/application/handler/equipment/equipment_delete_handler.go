package eq

import (
	"errors"
	eqcmd "octodome/internal/equipment/application/command/equipment"
	eqdom "octodome/internal/equipment/domain/equipment"
)

type DeleteHandler struct {
	validator eqdom.Validator
	repo      eqdom.Repository
}

func NewDeleteHandler(
	validator eqdom.Validator,
	repository eqdom.Repository) *DeleteHandler {
	return &DeleteHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *DeleteHandler) Handle(c eqcmd.Delete) error {
	if !h.validator.CanBeModified(c.ID, c.UserContext) {
		return errors.New("equipment cannot be deleted")
	}

	return h.repo.Delete(c.ID)
}
