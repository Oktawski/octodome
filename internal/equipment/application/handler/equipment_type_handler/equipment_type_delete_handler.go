package equipmenttype

import (
	"errors"
	eqcommand "octodome/internal/equipment/application/command"
	eqtypedom "octodome/internal/equipment/domain/equipment_type"
)

type DeleteHandler struct {
	validator eqtypedom.EquipmentTypeValidator
	repo      eqtypedom.EquipmentTypeRepository
}

func NewDeleteHandler(
	validator eqtypedom.EquipmentTypeValidator,
	repository eqtypedom.EquipmentTypeRepository) *DeleteHandler {

	return &DeleteHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *DeleteHandler) Handle(c eqcommand.EquipmentTypeDeleteCommand) error {

	if !h.validator.CanBeModified(c.ID, c.UserContext) {
		return errors.New("equipment type cannot be removed")
	}

	if err := h.repo.Delete(c.ID); err != nil {
		return err
	}

	return nil
}
