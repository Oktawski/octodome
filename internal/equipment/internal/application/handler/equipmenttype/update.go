package hdl

import (
	"errors"
	cmd "octodome/internal/equipment/internal/application/command"
	"octodome/internal/equipment/internal/dependencies"
	domain "octodome/internal/equipment/internal/domain/equipmenttype"
)

type UpdateHandler struct {
	repo      domain.Repository
	validator domain.Validator
}

func NewUpdateHandler(deps dependencies.EquipmentTypeContainer) *UpdateHandler {
	return &UpdateHandler{
		repo:      deps.Repository,
		validator: deps.Validator,
	}
}

func (h *UpdateHandler) Handle(c cmd.EquipmentTypeUpdate) error {
	if !h.validator.CanBeModified(c.ID, c.UserContext) {
		return errors.New("equipment type cannot be updated")
	}

	equipmentType := &domain.EquipmentType{
		ID:     c.ID,
		Name:   c.Name,
		UserID: c.UserContext.ID,
	}

	return h.repo.Update(equipmentType)
}
