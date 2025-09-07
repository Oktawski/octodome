package eq

import (
	"errors"
	eqcmd "octodome/internal/equipment/application/command/equipment"
	eqdom "octodome/internal/equipment/domain/equipment"
)

type CreateHandler struct {
	validator eqdom.Validator
	repo      eqdom.Repository
}

func NewCreateHandler(
	validator eqdom.Validator,
	repository eqdom.Repository) *CreateHandler {
	return &CreateHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *CreateHandler) Handle(c eqcmd.Create) error {
	if !h.validator.CanBeCreated(c.Name, c.EquipmentTypeID, c.UserContext) {
		return errors.New("equipment cannot be created")
	}

	equipment := &eqdom.Equipment{
		Name:            c.Name,
		EquipmentTypeID: c.EquipmentTypeID,
		UserID:          c.UserContext.ID,
	}

	return h.repo.Create(equipment)
}
