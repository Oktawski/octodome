package application

import (
	"context"
	"encoding/json"
	"fmt"

	"octodome.com/eventbroker/domain"
)

type Forward struct {
	eventRepository domain.EventRepository
	handlerRegistry domain.HandlerRegistry
	dispatcher      domain.Dispatcher
}

func NewForward(
	eventRepository domain.EventRepository,
	handlerRegistry domain.HandlerRegistry,
	dispatcher domain.Dispatcher,
) *Forward {
	return &Forward{
		eventRepository: eventRepository,
		handlerRegistry: handlerRegistry,
		dispatcher:      dispatcher,
	}
}

func (f *Forward) Execute(ctx context.Context, eventType string, payload any) error {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	event, err := f.eventRepository.Save(ctx, eventType, payloadJSON)
	if err != nil {
		return err
	}

	handlers, err := f.handlerRegistry.GetHandlers(ctx, eventType)
	if err != nil {
		return fmt.Errorf("fetch handlers for %s: %w", eventType, err)
	}

	for _, h := range handlers {
		f.dispatcher.Dispatch(ctx, event, h)
	}

	return nil
}
