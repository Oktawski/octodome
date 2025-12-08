package presentation

import "encoding/json"

type getEventResponse struct {
	ID      uint            `json:"id"`
	Payload json.RawMessage `json:"payload"`
}
