package eq

import (
	"errors"
	eqcmd "octodome/internal/equipment/application/command/equipment"
	eqdom "octodome/internal/equipment/domain/equipment"
)

type UpdateHandler struct {
	validator  eqdom.Validator
	repository eqdom.Repository
}

func NewUpdateHandler(
	validator eqdom.Validator,
	repository eqdom.Repository,
) *UpdateHandler {
	return &UpdateHandler{
		repository: repository,
		validator:  validator,
	}
}

func (h *UpdateHandler) Handle(c eqcmd.Update) error {
	if !h.validator.CanBeModified(c.ID, c.UserContext) {
		return errors.New("equipment cannot be modified")
	}

	equipment := &eqdom.Equipment{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Category:    c.Category,
		UserID:      c.UserContext.ID,
	}

	return h.repository.Update(equipment)
}
