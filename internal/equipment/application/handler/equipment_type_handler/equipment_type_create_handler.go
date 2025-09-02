package equipmenttype

import (
	"errors"
	eqcommand "octodome/internal/equipment/application/command"
	eqtypedom "octodome/internal/equipment/domain/equipment_type"
)

type CreateHandler struct {
	validator eqtypedom.EquipmentTypeValidator
	repo      eqtypedom.EquipmentTypeRepository
}

func NewCreateHandler(
	validator eqtypedom.EquipmentTypeValidator,
	repository eqtypedom.EquipmentTypeRepository) *CreateHandler {

	return &CreateHandler{
		validator: validator,
		repo:      repository,
	}
}

func (h *CreateHandler) Handle(c eqcommand.EquipmentTypeCreateCommand) error {
	if !h.validator.CanBeCreated(c.Name, c.UserContext) {
		return errors.New("equipment type with this name already exists")
	}

	eqType := &eqtypedom.EquipmentType{
		Name:   c.Name,
		UserID: c.UserContext.ID,
	}

	return h.repo.Create(eqType)
}
