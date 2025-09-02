package equipment

import (
	"errors"
	eqcommand "octodome/internal/equipment/application/command"
	equipmentdom "octodome/internal/equipment/domain/equipment"
)

type UpdateHandler struct {
	validator  equipmentdom.EquipmentValidator
	repository equipmentdom.EquipmentRepository
}

func NewUpdateHandler(
	validator equipmentdom.EquipmentValidator,
	repository equipmentdom.EquipmentRepository,
) *UpdateHandler {
	return &UpdateHandler{
		repository: repository,
		validator:  validator,
	}
}

func (h *UpdateHandler) Handle(c eqcommand.EquipmentUpdateCommand) error {
	if !h.validator.CanBeModified(c.ID, c.UserContext) {
		return errors.New("equipment cannot be modified")
	}

	equipment := &equipmentdom.Equipment{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Category:    c.Category,
		UserID:      c.UserContext.ID,
	}

	return h.repository.Update(equipment)
}
