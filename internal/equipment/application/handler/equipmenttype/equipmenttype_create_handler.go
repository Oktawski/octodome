package eqtype

import (
	"errors"
	eqtypecmd "octodome/internal/equipment/application/command/equipmenttype"
	eqtypedom "octodome/internal/equipment/domain/equipmenttype"
)

type CreateHandler struct {
	validator eqtypedom.Validator
	repo      eqtypedom.Repository
}

func NewCreateHandler(
	validator eqtypedom.Validator,
	repository eqtypedom.Repository) *CreateHandler {

	return &CreateHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *CreateHandler) Handle(c eqtypecmd.Create) error {
	if !h.validator.CanBeCreated(c.Name, c.UserContext) {
		return errors.New("equipment type with this name already exists")
	}

	eqType := &eqtypedom.EquipmentType{
		Name:   c.Name,
		UserID: c.UserContext.ID,
	}

	return h.repo.Create(eqType)
}
