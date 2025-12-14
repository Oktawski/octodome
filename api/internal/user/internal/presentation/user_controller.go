package http

import (
	"net/http"

	user "octodome.com/api/internal/user/internal/application"
	corehttp "octodome.com/shared/http"
)

type UserController struct {
	createHandler  *user.CreateHandler
	getByIDHandler *user.GetByIDHandler
}

func NewUserController(
	create *user.CreateHandler,
	getByID *user.GetByIDHandler) *UserController {
	return &UserController{
		createHandler:  create,
		getByIDHandler: getByID,
	}
}

func (ctrl *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := corehttp.GetPathParam[int](r, "id")
	if err != nil || id < 0 {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	query := user.GetByID{Context: r.Context(), ID: uint(id)}
	user, err := ctrl.getByIDHandler.Handle(query)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusNotFound, "User not found")
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, user)
}

func (ctrl *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var command user.Create
	if err := corehttp.ParseJSON(r, &command); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ctrl.createHandler.Handle(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	command.Context = r.Context()

	// TODO: extend by ID
	corehttp.WriteJSON(w, http.StatusCreated, command)
}
