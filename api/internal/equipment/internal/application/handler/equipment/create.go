package hdl

import (
	"context"
	"errors"

	corecontext "octodome.com/api/internal/core/context"
	cmd "octodome.com/api/internal/equipment/internal/application/command"
	"octodome.com/api/internal/equipment/internal/dependencies"
	domain "octodome.com/api/internal/equipment/internal/domain/equipment"
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
	ctx := context.WithValue(c.Ctx, corecontext.UserIDKey, c.UserContext.ID)

	if !h.validator.CanBeCreated(ctx, c.Name, c.EquipmentTypeID) {
		return errors.New("equipment cannot be created")
	}

	equipment := &domain.Equipment{
		Name:            c.Name,
		EquipmentTypeID: c.EquipmentTypeID,
		UserID:          c.UserContext.ID,
	}

	return h.repo.Create(ctx, equipment)
}
