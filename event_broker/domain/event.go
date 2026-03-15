package domain

import "encoding/json"

type Event struct {
	ID      uint
	Type    string
	Payload json.RawMessage
}
