package application

import (
	"context"

	"octodome.com/eventbroker/domain"
)

type GetEvent struct {
	eventRepository domain.EventRepository
}

func NewGetEvent(eventRepository domain.EventRepository) *GetEvent {
	return &GetEvent{eventRepository: eventRepository}
}

func (g *GetEvent) Execute(ctx context.Context, eventType string) (domain.Event, error) {
	return g.eventRepository.Get(ctx, eventType)
}
