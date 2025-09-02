package equipment

import (
	"errors"
	eqcommand "octodome/internal/equipment/application/command"
	equipmentdom "octodome/internal/equipment/domain/equipment"
)

type DeleteHandler struct {
	validator equipmentdom.EquipmentValidator
	repo      equipmentdom.EquipmentRepository
}

func NewDeleteHandler(
	validator equipmentdom.EquipmentValidator,
	repository equipmentdom.EquipmentRepository) *DeleteHandler {
	return &DeleteHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *DeleteHandler) Handle(c eqcommand.EquipmentDeleteCommand) error {
	if !h.validator.CanBeModified(c.ID, c.UserContext) {
		return errors.New("equipment cannot be deleted")
	}

	return h.repo.Delete(c.ID)
}
