package infra

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"octodome.com/eventbroker/domain"
)

type dispatchRequest struct {
	ID      uint            `json:"id"`
	Payload json.RawMessage `json:"payload"`
}

type dispatcher struct {
	client *http.Client
}

func NewEventDispatcher() domain.Dispatcher {
	return &dispatcher{
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (d *dispatcher) Dispatch(ctx context.Context, event domain.Event, handler domain.Handler) error {
	body, _ := json.Marshal(dispatchRequest{ID: event.ID, Payload: event.Payload})
	resp, err := d.client.Post(handler.URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Failed to dispatch event %d to handler %s (%s): %v", event.ID, handler.Name, handler.URL, err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Printf("Handler %s returned non-2xx status %d for event %d", handler.Name, resp.StatusCode, event.ID)
	}

	return nil
}
