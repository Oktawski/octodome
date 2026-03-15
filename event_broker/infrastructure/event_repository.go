package infra

import (
	"context"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"octodome.com/eventbroker/domain"
	"octodome.com/eventbroker/infrastructure/model"
	"octodome.com/shared/events"
)

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) domain.EventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) Save(ctx context.Context, eventType string, payload []byte) (domain.Event, error) {
	eventModel := &model.Event{
		Type:      eventType,
		Payload:   payload,
		CreatedAt: time.Now(),
		Status:    string(events.EventStatusPending),
	}

	if err := gorm.G[model.Event](r.db).Create(ctx, eventModel); err != nil {
		return domain.Event{}, err
	}

	return domain.Event{
		ID:      eventModel.ID,
		Type:    eventModel.Type,
		Payload: eventModel.Payload,
	}, nil
}

func (r *eventRepository) GetStale(ctx context.Context) ([]domain.Event, error) {
	stale, err := gorm.G[model.Event](r.db, clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
		Where("status IN ? OR (status = 'processing' AND updated_at < NOW() - INTERVAL '5 minutes')",
			[]string{
				string(events.EventStatusPending),
				string(events.EventStatusFailed),
			}).
		Order("created_at ASC").
		Find(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.Event, len(stale))
	for i, e := range stale {
		result[i] = domain.Event{
			ID:      e.ID,
			Type:    e.Type,
			Payload: e.Payload,
		}
	}
	return result, nil
}

func (r *eventRepository) Get(ctx context.Context, eventType string) (domain.Event, error) {
	e, err := gorm.G[model.Event](r.db, clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
		Where("type = ? AND status IN ?",
			eventType,
			[]string{
				string(events.EventStatusPending),
				string(events.EventStatusFailed),
			}).
		Order("CASE status WHEN 'pending' THEN 1 WHEN 'failed' THEN 2 END, created_at ASC").
		First(ctx)
	if err != nil {
		return domain.Event{}, err
	}

	return domain.Event{
		ID:      e.ID,
		Type:    e.Type,
		Payload: e.Payload,
	}, nil
}
