package dependencies

import domain "octodome/internal/equipment/internal/domain/equipmenttype"

type EquipmentTypeContainer struct {
	Repository domain.Repository
	Validator  domain.Validator
}

func NewEquipmentTypeContainer(
	repository domain.Repository,
	validator domain.Validator,
) EquipmentTypeContainer {
	return EquipmentTypeContainer{
		Repository: repository,
		Validator:  validator,
	}
}
