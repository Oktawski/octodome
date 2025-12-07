package http

import (
	"net/http"
	"octodome.com/api/internal/auth/domain"
	auth "octodome.com/api/internal/auth/internal/application"
	"octodome.com/shared/collection"
	corehttp "octodome.com/api/internal/core/http"
)

type AuthController struct {
	AuthenticateHandler auth.AuthenticateHandler
	AssignRoleHandler   auth.AssignRoleHandler
	UnassignRoleHandler auth.UnassignRoleHandler
	SyncRolesHandler    auth.SyncRolesHandler
}

func NewAuthController(
	authenticateHandler auth.AuthenticateHandler,
	assignRoleHandler auth.AssignRoleHandler,
	unassignRoleHandler auth.UnassignRoleHandler,
	syncRolesHandler auth.SyncRolesHandler,
) *AuthController {
	return &AuthController{
		AuthenticateHandler: authenticateHandler,
		AssignRoleHandler:   assignRoleHandler,
		UnassignRoleHandler: unassignRoleHandler,
		SyncRolesHandler:    syncRolesHandler,
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
	userContext, _ := corehttp.GetUserContext(r)

	var request AssignRoleRequest
	if err := corehttp.ParseJSON(r, &request); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	cmd := auth.AssignRoleCommand{
		Role:        domain.RoleName(request.Role),
		UserID:      request.UserID,
		UserContext: *userContext,
	}

	err := ctrl.AssignRoleHandler.Handle(cmd)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, nil)
}

func (ctrl *AuthController) UnassignRole(w http.ResponseWriter, r *http.Request) {
	userContext, _ := corehttp.GetUserContext(r)

	var request UnassignRoleRequest
	if err := corehttp.ParseJSON(r, &request); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	cmd := auth.UnassignRoleCommand{
		Role:        domain.RoleName(request.Role),
		UserID:      request.UserID,
		UserContext: *userContext,
	}

	err := ctrl.UnassignRoleHandler.Handle(cmd)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, nil)
}

func (ctrl *AuthController) SyncRoles(w http.ResponseWriter, r *http.Request) {
	userContext, _ := corehttp.GetUserContext(r)

	var request SyncRolesRequest
	if err := corehttp.ParseJSON(r, &request); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	cmd := auth.SyncRolesCommand{
		Roles: collection.Map(request.Roles, func(e string) domain.RoleName {
			return domain.RoleName(e)
		}),
		UserID:      request.UserID,
		UserContext: *userContext,
	}

	err := ctrl.SyncRolesHandler.Handle(cmd)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusOK, nil)
}
