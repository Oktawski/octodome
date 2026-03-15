package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"octodome.com/shared/events"
	corehttp "octodome.com/shared/http"
)

type eventRequest struct {
	ID      uint            `json:"id"`
	Payload json.RawMessage `json:"payload"`
}

func newRouter(eventClient events.Client) http.Handler {
	r := chi.NewRouter()
	r.Post("/events", handleEvent(eventClient))
	return r
}

func handleEvent(eventClient events.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req eventRequest
		if err := corehttp.ParseJSON(r, &req); err != nil {
			corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		if err := eventClient.MarkEventAsProcessing(req.ID); err != nil {
			log.Printf("Failed to mark event %d as processing: %v", req.ID, err)
			corehttp.WriteJSONError(w, http.StatusInternalServerError, "failed to accept event")
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

		if _, err := sendEmail(payload); err != nil {
			log.Printf("Failed to process event %d: %v", req.ID, err)
			if err := eventClient.MarkEventAsFailed(req.ID); err != nil {
				log.Printf("Failed to mark event %d as failed: %v", req.ID, err)
			}
			corehttp.WriteJSONError(w, http.StatusInternalServerError, "failed to process event")
			return
		}

		if err := eventClient.MarkEventAsProcessed(req.ID); err != nil {
			log.Printf("Failed to mark event %d as processed: %v", req.ID, err)
		}

		corehttp.WriteJSON(w, http.StatusOK, nil)
	}
}
