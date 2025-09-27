package hdl

import (
	"errors"
	cmd "octodome/internal/equipment/internal/application/command"
	"octodome/internal/equipment/internal/dependencies"
	domain "octodome/internal/equipment/internal/domain/equipment"
)

type CreateHandler struct {
	repo      domain.Repository
	validator domain.Validator
}

func NewCreateHandler(deps dependencies.EquipmentContainer) *CreateHandler {
	return &CreateHandler{
		repo:      deps.Repository,
		validator: deps.Validator,
	}
}

func (h *CreateHandler) Handle(c cmd.EquipmentCreate) error {
	if !h.validator.CanBeCreated(c.Name, c.EquipmentTypeID, c.UserContext) {
		return errors.New("equipment cannot be created")
	}

	equipment := &domain.Equipment{
		Name:            c.Name,
		EquipmentTypeID: c.EquipmentTypeID,
		UserID:          c.UserContext.ID,
	}

	return h.repo.Create(equipment)
}
