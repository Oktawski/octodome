package http

import (
	"net/http"
	"octodome/internal/auth/domain"
	auth "octodome/internal/auth/internal/application"
	corehttp "octodome/internal/core/http"
)

type AuthController struct {
	AuthenticateHandler auth.AuthenticateHandler
	AssignRoleHandler   auth.AssignRoleHandler
}

func NewAuthController(
	authenticateHandler auth.AuthenticateHandler,
	assignRoleHandler auth.AssignRoleHandler,
) *AuthController {
	return &AuthController{
		AuthenticateHandler: authenticateHandler,
		AssignRoleHandler:   assignRoleHandler,
	}
}

func (ctrl *AuthController) Authenticate(w http.ResponseWriter, r *http.Request) {
	var request AuthRequest

	if err := corehttp.ParseJSON(r, &request); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	authCommand := auth.AuthenticateCommand{
		Username: request.Username,
		Password: request.Password,
	}

	token, err := ctrl.AuthenticateHandler.Handle(authCommand)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, token)
}

func (ctrl *AuthController) AssignRole(w http.ResponseWriter, r *http.Request) {
	var request AssignRoleRequest

	if err := corehttp.ParseJSON(r, &request); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	cmd := auth.AssignRoleCommand{
		Role:   domain.RoleName(request.Role),
		UserID: request.UserID,
	}

	err := ctrl.AssignRoleHandler.Handle(cmd)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, nil)
}
