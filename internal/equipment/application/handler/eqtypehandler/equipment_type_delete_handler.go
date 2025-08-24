package eqtypehandler

import (
	"errors"
	eqcommand "octodome/internal/equipment/application/command"
	eqdom "octodome/internal/equipment/domain"
)

type DeleteHandler struct {
	validator eqdom.EquipmentTypeValidator
	repo      eqdom.EquipmentTypeRepository
}

func NewDeleteHandler(
	validator eqdom.EquipmentTypeValidator,
	repository eqdom.EquipmentTypeRepository) *DeleteHandler {

	return &DeleteHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *DeleteHandler) Handle(c eqcommand.DeleteCommand) error {

	if !h.validator.CanBeDeleted(c.ID, c.User) {
		return errors.New("equipment type cannot be removed")
	}

	if err := h.repo.DeleteEquipmentType(c.ID, c.User); err != nil {
		return err
	}

	return nil
}
