package presentation

import (
	"encoding/json"
	"net/http"

	"octodome.com/eventbroker/domain"
	"octodome.com/shared/events"
	corehttp "octodome.com/shared/http"
)

type eventRequest struct {
	EventType events.EventType `json:"type"`
	Payload   json.RawMessage  `json:"payload"`
}

type EventController struct {
	publisher domain.Publisher
	consumer  domain.Consumer
}

func NewEventController(
	publisher domain.Publisher,
	consumer domain.Consumer,
) *EventController {
	return &EventController{
		publisher: publisher,
		consumer:  consumer,
	}
}

func (c *EventController) PublishEvent(w http.ResponseWriter, r *http.Request) {
	var request eventRequest
	if err := corehttp.ParseJSON(r, &request); err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	err := c.publisher.Publish(request.EventType, request.Payload)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusInternalServerError, "failed to publish event")
		return
	}

	corehttp.WriteJSON(w, http.StatusCreated, request)
}

func (c *EventController) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventType, err := corehttp.GetPathParam[string](r, "eventType")
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "event type is required")
		return
	}

	id, payload, err := c.consumer.GetEvent(eventType)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusNotFound, "no events found")
		return
	}

	response := getEventResponse{
		ID:      id,
		Payload: payload.(json.RawMessage),
	}

	corehttp.WriteJSON(w, http.StatusOK, response)
}

func (c *EventController) MarkEventAsPending(w http.ResponseWriter, r *http.Request) {
	id, err := corehttp.GetID(r)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "event ID is required")
		return
	}

	err = c.consumer.MarkEventAsPending(id)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusNoContent, nil)
}

func (c *EventController) MarkEventAsProcessing(w http.ResponseWriter, r *http.Request) {
	id, err := corehttp.GetID(r)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "event ID is required")
		return
	}

	err = c.consumer.MarkEventAsProcessing(id)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusNoContent, nil)
}

func (c *EventController) MarkEventAsProcessed(w http.ResponseWriter, r *http.Request) {
	id, err := corehttp.GetID(r)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "event ID is required")
		return
	}

	err = c.consumer.MarkEventAsProcessed(id)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusNoContent, nil)
}

func (c *EventController) MarkEventAsFailed(w http.ResponseWriter, r *http.Request) {
	id, err := corehttp.GetID(r)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusBadRequest, "event ID is required")
		return
	}

	err = c.consumer.MarkEventAsFailed(id)
	if err != nil {
		corehttp.WriteJSONError(w, http.StatusConflict, err.Error())
		return
	}

	corehttp.WriteJSON(w, http.StatusNoContent, nil)
}
