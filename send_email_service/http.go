package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"octodome.com/send_email_service/handler"
	"octodome.com/shared/events"
	corehttp "octodome.com/shared/http"
)

type eventRequest struct {
	ID      uint            `json:"id"`
	Payload json.RawMessage `json:"payload"`
}

func newRouter(eventClient events.Client) http.Handler {
	r := chi.NewRouter()

	userRegisteredHandler := handler.NewUserRegisteredHandler(eventClient)
	magicCodeRequestedHandler := handler.NewMagicCodeRequestedHandler(eventClient)

	r.Post("/events/user-registered", handleUserRegisteredEvent(userRegisteredHandler, eventClient))
	r.Post("/events/magic-code-requested", handleMagicCodeRequestedEvent(magicCodeRequestedHandler, eventClient))
	return r
}

func handleUserRegisteredEvent(
	userRegisteredHandler handler.UserRegisteredHandler,
	eventClient events.Client,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req eventRequest
		if err := corehttp.ParseJSON(r, &req); err != nil {
			corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var payload events.UserRegistered
		if err := json.Unmarshal(req.Payload, &payload); err != nil {
			log.Printf("Failed to unmarshal event %d: %v", req.ID, err)
			if err := eventClient.MarkEventAsFailed(req.ID); err != nil {
				log.Printf("Failed to mark event %d as failed: %v", req.ID, err)
			}
			corehttp.WriteJSONError(w, http.StatusUnprocessableEntity, "invalid event payload")
			return
		}

		if _, err := userRegisteredHandler.Handle(req.ID, payload); err != nil {
			log.Printf("Failed to process event %d: %v", req.ID, err)
			corehttp.WriteJSONError(w, http.StatusInternalServerError, "failed to process event")
			return
		}

		corehttp.WriteJSON(w, http.StatusOK, nil)
	}
}

func handleMagicCodeRequestedEvent(
	magicCodeRequestedHandler handler.MagicCodeRequestedHandler,
	eventClient events.Client,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req eventRequest
		if err := corehttp.ParseJSON(r, &req); err != nil {
			corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		var payload events.MagicCodeRequested
		if err := json.Unmarshal(req.Payload, &payload); err != nil {
			log.Printf("Failed to unmarshal event %d: %v", req.ID, err)
			if err := eventClient.MarkEventAsFailed(req.ID); err != nil {
				log.Printf("Failed to mark event %d as failed: %v", req.ID, err)
			}
			corehttp.WriteJSONError(w, http.StatusUnprocessableEntity, "invalid event payload")
			return
		}

		if _, err := magicCodeRequestedHandler.Handle(req.ID, payload); err != nil {
			log.Printf("Failed to process event %d: %v", req.ID, err)
			corehttp.WriteJSONError(w, http.StatusInternalServerError, "failed to process event")
			return
		}

		corehttp.WriteJSON(w, http.StatusOK, nil)
	}
}
