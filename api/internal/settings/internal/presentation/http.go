package presentation

import (
	"encoding/json"
	"net/http"

	corehttp "octodome.com/api/internal/core/http"
	settings "octodome.com/api/internal/settings/internal/application"
	sharedhttp "octodome.com/shared/http"
)

type HttpHandler struct {
	upsertHandler *settings.UpsertHandler
}

func NewHttpHandler(upsertHandler *settings.UpsertHandler) *HttpHandler {
	return &HttpHandler{upsertHandler: upsertHandler}
}

func (h *HttpHandler) UpsertSetting(w http.ResponseWriter, r *http.Request) {
	user, _ := corehttp.GetUserContext(r)

	var request UpsertRequest
	if err := sharedhttp.ParseJSON(r, &request); err != nil {
		sharedhttp.WriteJSONError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	command := settings.Upsert{
		Context:     r.Context(),
		UserContext: *user,
		Name:        request.Name,
		Value:       string(request.Value),
	}

	err := h.upsertHandler.Handle(command)
	if err != nil {
		sharedhttp.WriteJSONError(w, http.StatusInternalServerError, "failed to upsert setting")
		return
	}

	sharedhttp.WriteJSON(w, http.StatusOK, nil)
}

type UpsertRequest struct {
	Name  string          `json:"name"`
	Value json.RawMessage `json:"value"`
}
