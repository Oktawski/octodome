package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"octodome.com/eventbroker/application"
	"octodome.com/shared/events"
	corehttp "octodome.com/shared/http"
)

type eventRequest struct {
	EventType string          `json:"type"`
	Payload   json.RawMessage `json:"payload"`
}

type eventController struct {
	forward         *application.Forward
	updateState     *application.UpdateState
	getEvent        *application.GetEvent
	registerHandler *application.RegisterHandler
}

func NewEventController(
	forward *application.Forward,
	updateState *application.UpdateState,
	getEvent *application.GetEvent,
	registerHandler *application.RegisterHandler,
) *eventController {
	return &eventController{
		forward:         forward,
		updateState:     updateState,
		getEvent:        getEvent,
		registerHandler: registerHandler,
	}
}

func RegisterEventRoutes(
	r chi.Router,
	forward *application.Forward,
	updateState *application.UpdateState,
	getEvent *application.GetEvent,
	registerHandler *application.RegisterHandler,
) {

	controller := NewEventController(forward, updateState, getEvent, registerHandler)

	r.Route("/handlers", func(handlers chi.Router) {
		handlers.Post("/", controller.RegisterHandler)
	})

	r.Route("/events", func(e chi.Router) {
		e.Post("/", controller.PublishEvent)
		e.Get("/{eventType}", controller.GetEvent)
		e.Put("/{id:[0-9]+}/pending", func(w http.ResponseWriter, r *http.Request) {
			controller.updateEventState(w, r, events.EventStatusPending)
		})
		e.Put("/{id:[0-9]+}/processing", func(w http.ResponseWriter, r *http.Request) {
			controller.updateEventState(w, r, events.EventStatusProcessing)
		})
		e.Put("/{id:[0-9]+}/processed", func(w http.ResponseWriter, r *http.Request) {
			controller.updateEventState(w, r, events.EventStatusProcessed)
		})
		e.Put("/{id:[0-9]+}/failed", func(w http.ResponseWriter, r *http.Request) {
			controller.updateEventState(w, r, events.EventStatusFailed)
		})
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
}

func (c *eventController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req registerHandlerRequest
	if err := corehttp.ParseJSON(r, &req); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	if req.Name == "" || req.EventType == "" || req.URL == "" {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "name, event_type and url are required")
		return
	}

	if err := c.registerHandler.Execute(r.Context(), req.Name, req.EventType, req.URL); err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, "failed to register handler")
		return
	}

	corehttp.WriteJSON(w, http.StatusCreated, nil)
}

func (c *eventController) PublishEvent(w http.ResponseWriter, r *http.Request) {
	var request eventRequest
	if err := corehttp.ParseJSON(r, &request); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	err := c.forward.Execute(r.Context(), request.EventType, request.Payload)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, "failed to publish event")
		return
	}

	corehttp.WriteJSON(w, http.StatusCreated, request)
}

func (c *eventController) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventType, err := corehttp.GetPathParam[string](r, "eventType")
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "event type is required")
		return
	}

	event, err := c.getEvent.Execute(r.Context(), eventType)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusNotFound, "no events found")
		return
	}

	response := getEventResponse{
		ID:      event.ID,
		Payload: event.Payload,
	}

	corehttp.WriteJSON(w, http.StatusOK, response)
}

func (c *eventController) updateEventState(w http.ResponseWriter, r *http.Request, state events.EventStatus) {
	id, err := corehttp.GetID(r)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "event ID is required")
		return
	}

	if err := c.updateState.Handle(r.Context(), id, state); err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusNoContent, nil)
}
