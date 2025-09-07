package eqtype

import (
	"errors"
	eqtypecmd "octodome/internal/equipment/application/command/equipmenttype"
	eqtypedom "octodome/internal/equipment/domain/equipmenttype"
)

type DeleteHandler struct {
	validator eqtypedom.Validator
	repo      eqtypedom.Repository
}

func NewDeleteHandler(
	validator eqtypedom.Validator,
	repository eqtypedom.Repository) *DeleteHandler {

	return &DeleteHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *DeleteHandler) Handle(c eqtypecmd.Delete) error {

	if !h.validator.CanBeModified(c.ID, c.UserContext) {
		return errors.New("equipment type cannot be removed")
	}

	if err := h.repo.Delete(c.ID); err != nil {
		return err
	}

	return nil
}
