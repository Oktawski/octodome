package hdl

import (
	"context"
	"errors"

	corecontext "octodome.com/api/internal/core/context"
	cmd "octodome.com/api/internal/equipment/internal/application/command"
	"octodome.com/api/internal/equipment/internal/dependencies"
	domain "octodome.com/api/internal/equipment/internal/domain/equipment"
)

type UpdateHandler struct {
	repository domain.Repository
	validator  domain.Validator
}

func NewUpdateHandler(deps dependencies.EquipmentContainer) *UpdateHandler {
	return &UpdateHandler{
		repository: deps.Repository,
		validator:  deps.Validator,
	}
}

func (h *UpdateHandler) Handle(c cmd.EquipmentUpdate) error {
	ctx := context.WithValue(c.Ctx, corecontext.UserIDKey, c.UserContext.ID)

	if !h.validator.CanBeModified(ctx, c.ID) {
		return errors.New("equipment cannot be modified")
	}

	equipment := &domain.Equipment{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Category:    c.Category,
		UserID:      c.UserContext.ID,
	}

	return h.repository.Update(ctx, equipment)
}
