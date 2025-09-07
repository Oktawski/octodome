package hdl

import (
	"errors"
	cmd "octodome/internal/equipment/application/command"
	domain "octodome/internal/equipment/domain/equipmenttype"
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

func (h *DeleteHandler) Handle(c cmd.EquipmentTypeDelete) error {

	if !h.validator.CanBeModified(c.ID, c.UserContext) {
		return errors.New("equipment type cannot be removed")
	}

	if err := h.repo.Delete(c.ID); err != nil {
		return err
	}

	return nil
}
