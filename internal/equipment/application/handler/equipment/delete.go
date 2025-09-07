package hdl

import (
	"errors"
	cmd "octodome/internal/equipment/application/command"
	domain "octodome/internal/equipment/domain/equipment"
)

type DeleteHandler struct {
	validator domain.Validator
	repo      domain.Repository
}

func NewDeleteHandler(
	validator domain.Validator,
	repository domain.Repository) *DeleteHandler {
	return &DeleteHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *DeleteHandler) Handle(c cmd.EquipmentDelete) error {
	if !h.validator.CanBeModified(c.ID, c.UserContext) {
		return errors.New("equipment cannot be deleted")
	}

	return h.repo.Delete(c.ID)
}
