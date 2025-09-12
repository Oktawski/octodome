package http

import (
	"net/http"
	auth "octodome/internal/auth/internal/application"
	corehttp "octodome/internal/core/http"
)

type AuthController struct {
	AuthenticateHandler auth.AuthenticateHandler
}

func NewAuthController(handler auth.AuthenticateHandler) *AuthController {
	return &AuthController{AuthenticateHandler: handler}
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

	token, err := ctrl.AuthenticateHandler.Handle(&authCommand)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, token)
}
