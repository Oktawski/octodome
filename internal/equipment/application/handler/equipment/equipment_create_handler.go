package equipment

import (
	"errors"
	eqcommand "octodome/internal/equipment/application/command"
	equipmentdom "octodome/internal/equipment/domain/equipment"
)

type CreateHandler struct {
	validator equipmentdom.EquipmentValidator
	repo      equipmentdom.EquipmentRepository
}

func NewCreateHandler(
	validator equipmentdom.EquipmentValidator,
	repository equipmentdom.EquipmentRepository) *CreateHandler {
	return &CreateHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *CreateHandler) Handle(c eqcommand.EquipmentCreateCommand) error {
	if !h.validator.CanBeCreated(c.Name, c.EquipmentTypeID, c.UserContext) {
		return errors.New("equipment cannot be created")
	}

	equipment := &equipmentdom.Equipment{
		Name:            c.Name,
		EquipmentTypeID: c.EquipmentTypeID,
		UserID:          c.UserContext.ID,
	}

	return h.repo.Create(equipment)
}
