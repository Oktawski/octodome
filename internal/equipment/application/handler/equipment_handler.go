package eqhandler

type EquipmentHandler interface {
}

type equipmentHandler struct {
}

func NewEquipmentHandler() *equipmentHandler {
	return &equipmentHandler{}
}
