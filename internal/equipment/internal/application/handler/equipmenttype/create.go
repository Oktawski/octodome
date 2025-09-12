package hdl

import (
	"errors"
	cmd "octodome/internal/equipment/internal/application/command"
	domain "octodome/internal/equipment/internal/domain/equipmenttype"
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

func (h *CreateHandler) Handle(c cmd.EquipmentTypeCreate) error {
	if !h.validator.CanBeCreated(c.Name, c.UserContext) {
		return errors.New("equipment type with this name already exists")
	}

	eqType := &domain.EquipmentType{
		Name:   c.Name,
		UserID: c.UserContext.ID,
	}

	return h.repo.Create(eqType)
}
