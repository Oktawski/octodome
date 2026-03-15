package application

import (
	"context"
	"fmt"

	"octodome.com/eventbroker/domain"
	"octodome.com/shared/events"
)

type UpdateState struct {
	stateManager domain.StateManager
	strategies   map[events.EventStatus]func(context.Context, uint) error
}

func NewUpdateState(stateManager domain.StateManager) *UpdateState {
	u := &UpdateState{stateManager: stateManager}
	u.strategies = map[events.EventStatus]func(context.Context, uint) error{
		events.EventStatusPending:    stateManager.MarkEventAsPending,
		events.EventStatusProcessing: stateManager.MarkEventAsProcessing,
		events.EventStatusProcessed:  stateManager.MarkEventAsProcessed,
		events.EventStatusFailed:     stateManager.MarkEventAsFailed,
	}
	return u
}

func (u *UpdateState) Handle(ctx context.Context, id uint, state events.EventStatus) error {
	strategy, ok := u.strategies[state]
	if !ok {
		return fmt.Errorf("unknown event state: %s", state)
	}
	return strategy(ctx, id)
}
