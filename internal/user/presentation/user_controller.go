package userpres

import (
	"net/http"
	corehttp "octodome/internal/core/http"
	usercommand "octodome/internal/user/application/command"
	userhandler "octodome/internal/user/application/handler"
	userquery "octodome/internal/user/application/query"
)

type UserController struct {
	userCreateHandler  *userhandler.CreateHandler
	userGetByIDHandler *userhandler.GetByIDHandler
}

func NewUserController(
	userCreate *userhandler.CreateHandler,
	userGetByID *userhandler.GetByIDHandler) *UserController {
	return &UserController{
		userCreateHandler:  userCreate,
		userGetByIDHandler: userGetByID,
	}
}

func (ctrl *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := corehttp.GetPathParam[int](r, "id")
	if err != nil || id < 0 {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	query := userquery.GetByID{ID: uint(id)}
	user, err := ctrl.userGetByIDHandler.Handle(query)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusNotFound, "User not found")
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, user)
}

func (ctrl *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var command usercommand.Create
	if err := corehttp.ParseJSON(r, &command); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := ctrl.userCreateHandler.Handle(command); err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: extend by ID
	corehttp.WriteJSON(w, http.StatusCreated, command)
}
