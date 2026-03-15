package domain

import "context"

type Handler struct {
	ID        uint
	Name      string
	EventType string
	URL       string
}

type HandlerRegistry interface {
	Register(ctx context.Context, name, eventType, url string) error
	GetHandlers(ctx context.Context, eventType string) ([]Handler, error)
}
