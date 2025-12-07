package hdl

import (
	"errors"
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
	if !h.validator.CanBeModified(c.ID, c.UserContext) {
		return errors.New("equipment cannot be deleted")
	}

	return h.repo.Delete(c.ID)
}
