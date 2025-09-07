package eqtype

import (
	"errors"
	eqtypecmd "octodome/internal/equipment/application/command/equipmenttype"
	eqtypedom "octodome/internal/equipment/domain/equipmenttype"
)

type UpdateHandler struct {
	validator eqtypedom.Validator
	repo      eqtypedom.Repository
}

func NewUpdateHandler(
	validator eqtypedom.Validator,
	repo eqtypedom.Repository,
) *UpdateHandler {
	return &UpdateHandler{
		validator: validator,
		repo:      repo,
	}
}

func (h *UpdateHandler) Handle(c eqtypecmd.Update) error {
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
