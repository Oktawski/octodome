package eqhttp

import (
	"net/http"
	corehttp "octodome/internal/core/http"
	eqcmd "octodome/internal/equipment/application/command/equipment"
	eq "octodome/internal/equipment/application/handler/equipment"
	eqqry "octodome/internal/equipment/application/query/equipment"
)

type EquipmentController struct {
	createHandler  *eq.CreateHandler
	updateHandler  *eq.UpdateHandler
	deleteHandler  *eq.DeleteHandler
	getByIDHandler *eq.GetByIDHandler
	getListHandler *eq.GetListHandler
}

func NewEquipmentController(
	createHandler *eq.CreateHandler,
	updateHandler *eq.UpdateHandler,
	deleteHandler *eq.DeleteHandler,
	getByIDHandler *eq.GetByIDHandler,
	getListHandler *eq.GetListHandler,
) *EquipmentController {
	return &EquipmentController{
		createHandler:  createHandler,
		updateHandler:  updateHandler,
		deleteHandler:  deleteHandler,
		getByIDHandler: getByIDHandler,
		getListHandler: getListHandler,
	}
}

func (c *EquipmentController) GetEquipmentList(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	pagination := corehttp.GetPagination(r)

	query := eqqry.GetList{Pagination: pagination, User: *user}

	equipments, totalCount, err := c.getListHandler.Handle(query)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusNotFound, "could not fetch equipments")
		return
	}

	response := &GetListResponse{Equipments: equipments, TotalCount: totalCount}
	corehttp.WriteJSON(w, http.StatusOK, response)
}

func (c *EquipmentController) GetEquipment(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	equipmentID, err := corehttp.GetID(r)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid equipment ID")
		return
	}

	query := eqqry.GetByID{ID: equipmentID, User: *user}

	equipment, err := c.getByIDHandler.Handle(query)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusNotFound, "could not fetch equipment")
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, equipment)
}

func (c *EquipmentController) CreateEquipment(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)

	var request CreateRequest
	if err := corehttp.ParseJSON(r, &request); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	command := eqcmd.Create{
		Name:        request.Name,
		Description: request.Description,
		Category:    request.Category,
		UserContext: *user,
	}
	if err := c.createHandler.Handle(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, "could not create equipment")
		return
	}

	corehttp.WriteJSON(w, http.StatusCreated, nil)
}

func (c *EquipmentController) UpdateEquipment(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	equipmentID, err := corehttp.GetID(r)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid equipment ID")
		return
	}

	var dto UpdateRequest
	if err := corehttp.ParseJSON(r, &dto); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	command := eqcmd.Update{
		ID:          equipmentID,
		Name:        dto.Name,
		Description: dto.Description,
		Category:    dto.Category,
		UserContext: *user,
	}
	if err := c.updateHandler.Handle(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, "could not update equipment")
		return
	}

	corehttp.WriteJSON(w, http.StatusNoContent, nil)
}

func (c *EquipmentController) DeleteEquipment(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	equipmentID, err := corehttp.GetID(r)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid equipment ID")
		return
	}

	command := eqcmd.Delete{
		ID:          equipmentID,
		UserContext: *user,
	}
	if err := c.deleteHandler.Handle(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, "could not delete equipment")
		return
	}

	corehttp.WriteJSON(w, http.StatusNoContent, nil)
}
