package hdl

import (
	"errors"
	cmd "octodome/internal/equipment/internal/application/command"
	domain "octodome/internal/equipment/internal/domain/equipmenttype"
)

type UpdateHandler struct {
	validator domain.Validator
	repo      domain.Repository
}

func NewUpdateHandler(
	validator domain.Validator,
	repo domain.Repository,
) *UpdateHandler {
	return &UpdateHandler{
		validator: validator,
		repo:      repo,
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
