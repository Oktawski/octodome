package corehttp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type HttpClient struct {
	baseURL string
	client  *http.Client
}

func NewHttpClient(baseURL string) *HttpClient {
	return &HttpClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (c *HttpClient) Get(url string, v any) error {
	response, err := c.client.Get(c.baseURL + url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, v); err != nil {
		return err
	}

	return nil
}

func (c *HttpClient) Post(url string, payload any, response *any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	httpResponse, err := c.client.Post(
		c.baseURL+url,
		"application/json",
		bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()

	data, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, response); err != nil {
		return err
	}

	return nil
}
