package equipmenttype

import (
	"errors"
	eqcommand "octodome/internal/equipment/application/command"
	eqtypedom "octodome/internal/equipment/domain/equipment_type"
)

type UpdateHandler struct {
	validator eqtypedom.EquipmentTypeValidator
	repo      eqtypedom.EquipmentTypeRepository
}

func NewUpdateHandler(
	validator eqtypedom.EquipmentTypeValidator,
	repo eqtypedom.EquipmentTypeRepository,
) *UpdateHandler {
	return &UpdateHandler{
		validator: validator,
		repo:      repo,
	}
}

func (h *UpdateHandler) Handle(c eqcommand.EquipmentTypeUpdateCommand) error {
	if !h.validator.CanBeModified(c.ID, c.UserContext) {
		return errors.New("equipment type cannot be updated")
	}

	equipmentType := &eqtypedom.EquipmentType{
		ID:     c.ID,
		Name:   c.Name,
		UserID: c.UserContext.ID,
	}

	return h.repo.Update(equipmentType)
}
