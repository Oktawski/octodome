package dependencies

import domain "octodome/internal/equipment/internal/domain/equipment"

type EquipmentContainer struct {
	Repository domain.Repository
	Validator  domain.Validator
}

func NewEquipmentContainer(
	repository domain.Repository,
	validator domain.Validator,
) EquipmentContainer {
	return EquipmentContainer{
		Repository: repository,
		Validator:  validator,
	}
}
