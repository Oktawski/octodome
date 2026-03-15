package presentation

import "encoding/json"

type getEventResponse struct {
	ID      uint            `json:"id"`
	Payload json.RawMessage `json:"payload"`
}

type registerHandlerRequest struct {
	Name      string `json:"name"`
	EventType string `json:"event_type"`
	URL       string `json:"url"`
}
