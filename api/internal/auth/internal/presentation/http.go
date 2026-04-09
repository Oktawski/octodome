package http

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
	"octodome.com/api/internal/auth/domain"
	auth "octodome.com/api/internal/auth/internal/application"
	corehttp "octodome.com/api/internal/core/http"
	"octodome.com/shared/collection"
	sharedhttp "octodome.com/shared/http"
)

type AuthHandlers struct {
	AuthenticateCredentialsHandler auth.AuthenticateCredentialsHandler
	SendMagicCodeHandler           auth.SendMagicCodeHandler
	AuthenticateMagicCodeHandler   auth.AuthenticateMagicCodeHandler
	AssignRoleHandler              auth.AssignRoleHandler
	UnassignRoleHandler            auth.UnassignRoleHandler
	SyncRolesHandler               auth.SyncRolesHandler
}

func NewAuthHandlers(
	authenticateCredentialsHandler auth.AuthenticateCredentialsHandler,
	sendMagicCodeHandler auth.SendMagicCodeHandler,
	authenticateMagicCodeHandler auth.AuthenticateMagicCodeHandler,
	assignRoleHandler auth.AssignRoleHandler,
	unassignRoleHandler auth.UnassignRoleHandler,
	syncRolesHandler auth.SyncRolesHandler,
) *AuthHandlers {
	return &AuthHandlers{
		AuthenticateCredentialsHandler: authenticateCredentialsHandler,
		SendMagicCodeHandler:           sendMagicCodeHandler,
		AuthenticateMagicCodeHandler:   authenticateMagicCodeHandler,
		AssignRoleHandler:              assignRoleHandler,
		UnassignRoleHandler:            unassignRoleHandler,
		SyncRolesHandler:               syncRolesHandler,
	}
}

func (handlers *AuthHandlers) AuthenticateCredentials(w http.ResponseWriter, r *http.Request) {
	var request AuthenticateCredentialsRequest

	if err := sharedhttp.ParseJSON(r, &request); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	authCommand := auth.AuthenticateCommand{
		Context:  r.Context(),
		Email:    request.Email,
		Password: request.Password,
	}

	token, err := handlers.AuthenticateCredentialsHandler.Handle(authCommand)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sharedhttp.WriteJSONError(w, http.StatusNotFound, "user not found")
			return
		}
		sharedhttp.WriteJSONError(w, http.StatusConflict, err.Error())
		return
	}

	response := AuthenticateResponse{
		AuthToken: token,
	}

	sharedhttp.WriteJSON(w, http.StatusOK, response)
}

func (handlers *AuthHandlers) AuthenticateMagicCode(w http.ResponseWriter, r *http.Request) {
	var request AuthenticateMagicCodeRequest
	if err := sharedhttp.ParseJSON(r, &request); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	authenticateMagicCodeCommand := auth.AuthenticateMagicCodeCommand{
		Context: r.Context(),
		Email:   request.Email,
		Code:    request.Code,
	}

	authToken, err := handlers.AuthenticateMagicCodeHandler.Handle(authenticateMagicCodeCommand)
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := AuthenticateResponse{
		AuthToken: authToken,
	}
	sharedhttp.WriteJSON(w, http.StatusOK, response)
}

func (handlers *AuthHandlers) SendMagicCode(w http.ResponseWriter, r *http.Request) {
	var request SendMagicCodeRequest

	if err := sharedhttp.ParseJSON(r, &request); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
	}

	sendMagicCodeCommand := auth.SendMagicCodeCommand{
		Context: r.Context(),
		Email:   request.Email,
	}

	err := handlers.SendMagicCodeHandler.Handle(sendMagicCodeCommand)
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sharedhttp.WriteJSON(w, http.StatusOK, nil)
}

func (handlers *AuthHandlers) AssignRole(w http.ResponseWriter, r *http.Request) {
	userContext, _ := corehttp.GetUserContext(r)

	var request AssignRoleRequest
	if err := sharedhttp.ParseJSON(r, &request); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	cmd := auth.AssignRoleCommand{
		Context:     r.Context(),
		Role:        domain.RoleName(request.Role),
		UserID:      request.UserID,
		UserContext: *userContext,
	}

	err := handlers.AssignRoleHandler.Handle(cmd)
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sharedhttp.WriteJSON(w, http.StatusOK, nil)
}

func (handlers *AuthHandlers) UnassignRole(w http.ResponseWriter, r *http.Request) {
	userContext, _ := corehttp.GetUserContext(r)

	var request UnassignRoleRequest
	if err := sharedhttp.ParseJSON(r, &request); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	cmd := auth.UnassignRoleCommand{
		Context:     r.Context(),
		Role:        domain.RoleName(request.Role),
		UserID:      request.UserID,
		UserContext: *userContext,
	}

	err := handlers.UnassignRoleHandler.Handle(cmd)
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sharedhttp.WriteJSON(w, http.StatusOK, nil)
}

func (handlers *AuthHandlers) SyncRoles(w http.ResponseWriter, r *http.Request) {
	userContext, _ := corehttp.GetUserContext(r)

	var request SyncRolesRequest
	if err := sharedhttp.ParseJSON(r, &request); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	cmd := auth.SyncRolesCommand{
		Context: r.Context(),
		Roles: collection.Map(request.Roles, func(e string) domain.RoleName {
			return domain.RoleName(e)
		}),
		UserID:      request.UserID,
		UserContext: *userContext,
	}

	err := handlers.SyncRolesHandler.Handle(cmd)
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sharedhttp.WriteJSON(w, http.StatusOK, nil)
}
