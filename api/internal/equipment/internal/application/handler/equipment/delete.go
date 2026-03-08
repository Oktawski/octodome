package hdl

import (
	"context"
	"errors"

	corecontext "octodome.com/api/internal/core/context"
	cmd "octodome.com/api/internal/equipment/internal/application/command"
	"octodome.com/api/internal/equipment/internal/dependencies"
	domain "octodome.com/api/internal/equipment/internal/domain/equipment"
)

type DeleteHandler struct {
	repo      domain.Repository
	validator domain.Validator
}

func NewDeleteHandler(deps dependencies.EquipmentContainer) *DeleteHandler {
	return &DeleteHandler{
		repo:      deps.Repository,
		validator: deps.Validator,
	}
}

func (h *DeleteHandler) Handle(c cmd.EquipmentDelete) error {
	ctx := context.WithValue(c.Ctx, corecontext.UserIDKey, c.UserContext.ID)

	if !h.validator.CanBeModified(ctx, c.ID) {
		return errors.New("equipment cannot be deleted")
	}

	return h.repo.Delete(ctx, c.ID)
}
