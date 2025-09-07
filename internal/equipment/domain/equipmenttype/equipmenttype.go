package eqtypedom

type EquipmentTypeDTO struct {
	ID   uint
	Name string
}

type EquipmentType struct {
	ID     uint
	Name   string
	UserID uint
}

func (et *EquipmentType) ToDTO() EquipmentTypeDTO {
	return EquipmentTypeDTO{
		ID:   et.ID,
		Name: et.Name,
	}
}
