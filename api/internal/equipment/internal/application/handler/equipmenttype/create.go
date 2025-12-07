package hdl

import (
	"errors"
	cmd "octodome.com/api/internal/equipment/internal/application/command"
	"octodome.com/api/internal/equipment/internal/dependencies"
	domain "octodome.com/api/internal/equipment/internal/domain/equipmenttype"
)

type CreateHandler struct {
	repo      domain.Repository
	validator domain.Validator
}

func NewCreateHandler(deps dependencies.EquipmentTypeContainer) *CreateHandler {

	return &CreateHandler{
		repo:      deps.Repository,
		validator: deps.Validator,
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
