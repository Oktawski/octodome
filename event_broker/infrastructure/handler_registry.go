package infra

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"octodome.com/eventbroker/domain"
	"octodome.com/eventbroker/infrastructure/model"
)

type handlerRegistry struct {
	db *gorm.DB
}

func NewHandlerRegistry(db *gorm.DB) domain.HandlerRegistry {
	return &handlerRegistry{db: db}
}

func (r *handlerRegistry) Register(ctx context.Context, name, eventType, url string) error {
	h := model.Handler{
		Name:      name,
		EventType: eventType,
		URL:       url,
	}

	return gorm.G[model.Handler](r.db, clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}, {Name: "event_type"}},
		DoUpdates: clause.AssignmentColumns([]string{"url", "updated_at"}),
	}).Create(ctx, &h)
}

func (r *handlerRegistry) GetHandlers(ctx context.Context, eventType string) ([]domain.Handler, error) {
	handlers, err := gorm.G[model.Handler](r.db).
		Where("event_type = ?", eventType).
		Find(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.Handler, len(handlers))
	for i, h := range handlers {
		result[i] = domain.Handler{
			ID:        h.ID,
			Name:      h.Name,
			EventType: h.EventType,
			URL:       h.URL,
		}
	}

	return result, nil
}
