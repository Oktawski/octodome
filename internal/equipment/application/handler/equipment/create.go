package hdl

import (
	"errors"
	cmd "octodome/internal/equipment/application/command"
	domain "octodome/internal/equipment/domain/equipment"
)

type CreateHandler struct {
	validator domain.Validator
	repo      domain.Repository
}

func NewCreateHandler(
	validator domain.Validator,
	repository domain.Repository) *CreateHandler {
	return &CreateHandler{
		validator: validator,
		repo:      repository,
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
