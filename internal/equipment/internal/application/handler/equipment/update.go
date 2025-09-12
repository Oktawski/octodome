package hdl

import (
	"errors"
	cmd "octodome/internal/equipment/internal/application/command"
	domain "octodome/internal/equipment/internal/domain/equipment"
)

type UpdateHandler struct {
	validator  domain.Validator
	repository domain.Repository
}

func NewUpdateHandler(
	validator domain.Validator,
	repository domain.Repository,
) *UpdateHandler {
	return &UpdateHandler{
		repository: repository,
		validator:  validator,
	}
}

func (h *UpdateHandler) Handle(c cmd.EquipmentUpdate) error {
	if !h.validator.CanBeModified(c.ID, c.UserContext) {
		return errors.New("equipment cannot be modified")
	}

	equipment := &domain.Equipment{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Category:    c.Category,
		UserID:      c.UserContext.ID,
	}

	return h.repository.Update(equipment)
}
