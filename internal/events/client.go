package events

import (
	"encoding/json"
	"fmt"
	"net/http"

	corehttp "octodome.com/shared/http"
)

type Client struct {
	client *corehttp.HttpClient
}

func NewClient(baseUrl string) *Client {
	return &Client{
		client: corehttp.NewHttpClient(baseUrl),
	}
}

type eventRequest struct {
	EventType EventType `json:"type"`
	Payload   any       `json:"payload"`
}

type getEventResponse struct {
	ID      uint            `json:"id"`
	Payload json.RawMessage `json:"payload"`
}

func (c *Client) PublishEvent(eventType EventType, payload EventType) error {
	request := eventRequest{
		EventType: eventType,
		Payload:   payload,
	}

	err := c.client.Post("", request, nil)
	return err
}

func (c *Client) GetEvent(eventType EventType) (uint, json.RawMessage, error) {
	var resp getEventResponse
	err := c.client.Get(fmt.Sprintf("/%s", eventType), &resp)
	if err != nil {
		return 0, nil, err
	}
	return resp.ID, resp.Payload, nil
}

func (c *Client) MarkEventAsFailed(id uint) error {
	statusCode, err := c.client.Put(fmt.Sprintf("/%d/failed", id), nil, nil)
	if err != nil {
		return err
	}
	if statusCode != http.StatusOK && statusCode != http.StatusNoContent {
		return fmt.Errorf("mark event as failed: unexpected status %d", statusCode)
	}
	return nil
}

func (c *Client) MarkEventAsProcessed(id uint) error {
	statusCode, err := c.client.Put(fmt.Sprintf("/%d/processed", id), nil, nil)
	if err != nil {
		return err
	}
	if statusCode != http.StatusOK && statusCode != http.StatusNoContent {
		return fmt.Errorf("mark event as processed: unexpected status %d", statusCode)
	}
	return nil
}
