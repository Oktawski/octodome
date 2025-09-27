package hdl

import (
	"errors"
	cmd "octodome/internal/equipment/internal/application/command"
	"octodome/internal/equipment/internal/dependencies"
	domain "octodome/internal/equipment/internal/domain/equipmenttype"
)

type DeleteHandler struct {
	repo      domain.Repository
	validator domain.Validator
}

func NewDeleteHandler(deps dependencies.EquipmentTypeContainer) *DeleteHandler {

	return &DeleteHandler{
		repo:      deps.Repository,
		validator: deps.Validator,
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
