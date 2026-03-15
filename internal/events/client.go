package events

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	corehttp "octodome.com/shared/http"
)

const (
	defaultRetryAttempts = 3
	defaultRetryDelay    = 500 * time.Millisecond
)

type Client struct {
	client        *corehttp.HttpClient
	retryAttempts int
	retryDelay    time.Duration
}

func NewClient(baseUrl string) *Client {
	return &Client{
		client:        corehttp.NewHttpClient(baseUrl),
		retryAttempts: defaultRetryAttempts,
		retryDelay:    defaultRetryDelay,
	}
}

func (c *Client) retry(fn func() error) error {
	var err error
	for range c.retryAttempts {
		if err = fn(); err == nil {
			return nil
		}
		time.Sleep(c.retryDelay)
	}
	return err
}

type registerHandlerRequest struct {
	Name      string    `json:"name"`
	EventType EventType `json:"event_type"`
	URL       string    `json:"url"`
}

type eventRequest struct {
	EventType EventType `json:"type"`
	Payload   any       `json:"payload"`
}

type getEventResponse struct {
	ID      uint            `json:"id"`
	Payload json.RawMessage `json:"payload"`
}

func (c *Client) RegisterHandler(name string, eventType EventType, url string) error {
	request := registerHandlerRequest{
		Name:      name,
		EventType: eventType,
		URL:       url,
	}
	return c.retry(func() error {
		return c.client.Post("/handlers/", request, nil)
	})
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
	err := c.client.Get(fmt.Sprintf("/events/%s", eventType), &resp)
	if err != nil {
		return 0, nil, err
	}
	return resp.ID, resp.Payload, nil
}

func (c *Client) MarkEventAsProcessing(id uint) error {
	return c.retry(func() error {
		statusCode, err := c.client.Put(fmt.Sprintf("/events/%d/processing", id), nil, nil)
		if err != nil {
			return err
		}
		if statusCode != http.StatusOK && statusCode != http.StatusNoContent {
			return fmt.Errorf("mark event as processing: unexpected status %d", statusCode)
		}
		return nil
	})
}

func (c *Client) MarkEventAsFailed(id uint) error {
	return c.retry(func() error {
		statusCode, err := c.client.Put(fmt.Sprintf("/events/%d/failed", id), nil, nil)
		if err != nil {
			return err
		}
		if statusCode != http.StatusOK && statusCode != http.StatusNoContent {
			return fmt.Errorf("mark event as failed: unexpected status %d", statusCode)
		}
		return nil
	})
}

func (c *Client) MarkEventAsProcessed(id uint) error {
	return c.retry(func() error {
		statusCode, err := c.client.Put(fmt.Sprintf("/events/%d/processed", id), nil, nil)
		if err != nil {
			return err
		}
		if statusCode != http.StatusOK && statusCode != http.StatusNoContent {
			return fmt.Errorf("mark event as processed: unexpected status %d", statusCode)
		}
		return nil
	})
}
