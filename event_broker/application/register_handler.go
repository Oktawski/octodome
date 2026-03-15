package application

import (
	"context"

	"octodome.com/eventbroker/domain"
)

type RegisterHandler struct {
	handlerRegistry domain.HandlerRegistry
}

func NewRegisterHandler(handlerRegistry domain.HandlerRegistry) *RegisterHandler {
	return &RegisterHandler{handlerRegistry: handlerRegistry}
}

func (r *RegisterHandler) Execute(ctx context.Context, name, eventType, url string) error {
	return r.handlerRegistry.Register(ctx, name, eventType, url)
}
