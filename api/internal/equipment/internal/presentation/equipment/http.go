package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	corehttp "octodome.com/api/internal/core/http"
	cmd "octodome.com/api/internal/equipment/internal/application/command"
	hdl "octodome.com/api/internal/equipment/internal/application/handler/equipment"
	qry "octodome.com/api/internal/equipment/internal/application/query"
	"octodome.com/api/internal/equipment/internal/dependencies"
	"octodome.com/api/internal/web/middleware"
	sharedhttp "octodome.com/shared/http"
)

type equipmentController struct {
	createHandler  *hdl.CreateHandler
	updateHandler  *hdl.UpdateHandler
	deleteHandler  *hdl.DeleteHandler
	getByIDHandler *hdl.GetByIDHandler
	getListHandler *hdl.GetListHandler
}

func RegisterEquipmentRoutes(
	r chi.Router,
	db *gorm.DB,
	deps dependencies.EquipmentContainer,
) {
	create := hdl.NewCreateHandler(deps)
	update := hdl.NewUpdateHandler(deps)
	delete := hdl.NewDeleteHandler(deps)
	getByID := hdl.NewGetByIDHandler(deps)
	getList := hdl.NewGetListHandler(deps)

	ctrl := &equipmentController{
		createHandler:  create,
		updateHandler:  update,
		deleteHandler:  delete,
		getByIDHandler: getByID,
		getListHandler: getList,
	}

	r.Route("/equipment", func(eq chi.Router) {
		eq.Use(middleware.JwtAuthMiddleware(db))

		eq.Get("/", ctrl.GetEquipmentList)
		eq.Get("/{id:[0-9]+}", ctrl.GetEquipment)

		eq.Post("/", ctrl.CreateEquipment)
		eq.Put("/{id:[0-9]+}", ctrl.UpdateEquipment)
		eq.Delete("/{id:[0-9]+}", ctrl.DeleteEquipment)
	})
}

func (c *equipmentController) GetEquipmentList(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	pagination := sharedhttp.GetPagination(r)

	query := qry.EquipmentGetList{Pagination: pagination, User: *user}

	equipments, totalCount, err := c.getListHandler.Handle(query)
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusNotFound, "could not fetch equipments")
		return
	}

	response := &GetListResponse{Equipments: equipments, TotalCount: totalCount}
	sharedhttp.WriteJSON(w, http.StatusOK, response)
}

func (c *equipmentController) GetEquipment(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	equipmentID, err := sharedhttp.GetID(r)
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "invalid equipment ID")
		return
	}

	query := qry.EquipmentGetByID{ID: equipmentID, User: *user}

	equipment, err := c.getByIDHandler.Handle(query)
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusNotFound, "could not fetch equipment")
		return
	}

	sharedhttp.WriteJSON(w, http.StatusOK, equipment)
}

func (c *equipmentController) CreateEquipment(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)

	var request CreateRequest
	if err := sharedhttp.ParseJSON(r, &request); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	command := cmd.EquipmentCreate{
		UserContext:     *user,
		Name:            request.Name,
		Description:     request.Description,
		Category:        request.Category,
		EquipmentTypeID: request.Type,
	}
	if err := c.createHandler.Handle(command); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusConflict, "could not create equipment")
		return
	}

	sharedhttp.WriteJSON(w, http.StatusCreated, nil)
}

func (c *equipmentController) UpdateEquipment(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	equipmentID, err := sharedhttp.GetID(r)
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "invalid equipment ID")
		return
	}

	var dto UpdateRequest
	if err := sharedhttp.ParseJSON(r, &dto); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	command := cmd.EquipmentUpdate{
		UserContext: *user,
		ID:          equipmentID,
		Name:        dto.Name,
		Description: dto.Description,
		Category:    dto.Category,
	}
	if err := c.updateHandler.Handle(command); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusConflict, "could not update equipment")
		return
	}

	sharedhttp.WriteJSON(w, http.StatusNoContent, nil)
}

func (c *equipmentController) DeleteEquipment(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)
	equipmentID, err := sharedhttp.GetID(r)
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "invalid equipment ID")
		return
	}

	command := cmd.EquipmentDelete{
		ID:          equipmentID,
		UserContext: *user,
	}
	if err := c.deleteHandler.Handle(command); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusConflict, "could not delete equipment")
		return
	}

	sharedhttp.WriteJSON(w, http.StatusNoContent, nil)
}
