package infra

import (
	"context"
	"log"

	"octodome.com/eventbroker/domain"
)

type sweeper struct {
	eventRepository domain.EventRepository
	handlerRegistry domain.HandlerRegistry
	dispatcher      domain.Dispatcher
}

func NewSweeper(
	eventRepository domain.EventRepository,
	handlerRegistry domain.HandlerRegistry,
	dispatcher domain.Dispatcher,
) domain.Sweeper {
	return &sweeper{
		eventRepository: eventRepository,
		handlerRegistry: handlerRegistry,
		dispatcher:      dispatcher,
	}
}

func (s *sweeper) Sweep(ctx context.Context) error {
	stale, err := s.eventRepository.GetStale(ctx)
	if err != nil {
		return err
	}

	for _, event := range stale {
		handlers, err := s.handlerRegistry.GetHandlers(ctx, event.Type)
		if err != nil {
			log.Printf("Failed to get handlers for event %d (type %s): %v", event.ID, event.Type, err)
			continue
		}

		for _, h := range handlers {
			if err := s.dispatcher.Dispatch(ctx, event, h); err != nil {
				log.Printf("Failed to dispatch event %d to handler %s: %v", event.ID, h.Name, err)
			}
		}
	}

	return nil
}
