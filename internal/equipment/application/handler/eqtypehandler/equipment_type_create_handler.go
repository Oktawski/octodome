package eqtypehandler

import (
	"errors"
	eqcommand "octodome/internal/equipment/application/command"
	eqdom "octodome/internal/equipment/domain"
)

type CreateHandler struct {
	validator eqdom.EquipmentTypeValidator
	repo      eqdom.EquipmentTypeRepository
}

func NewCreateHandler(
	validator eqdom.EquipmentTypeValidator,
	repository eqdom.EquipmentTypeRepository) *CreateHandler {

	return &CreateHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *CreateHandler) Handle(c eqcommand.CreateCommand) error {
	if !h.validator.CanBeCreated(c.Name, c.User) {
		return errors.New("equipment type with this name already exists")
	}

	eqType := &eqdom.EquipmentType{
		Name:   c.Name,
		UserID: c.User.ID,
	}

	return h.repo.CreateType(eqType)
}
