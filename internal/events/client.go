package events

import (
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

func (c *Client) PublishEvent(eventType EventType, payload EventType) error {
	request := eventRequest{
		EventType: eventType,
		Payload:   payload,
	}

	err := c.client.Post("", request, nil)
	return err
}
