package infra

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"octodome.com/eventbroker/domain"
	"octodome.com/shared/events"
)

type consumer struct {
	db *gorm.DB
}

func NewEventConsumer(db *gorm.DB) domain.Consumer {
	return &consumer{db: db}
}

func (c *consumer) GetEvent(ctx context.Context, eventType string) (uint, interface{}, error) {
	var eventModel event
	err := c.db.Transaction(func(tx *gorm.DB) error {
		e, err := gorm.G[event](tx, clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
			Where("type = ? AND status IN ?",
				eventType,
				[]string{
					string(events.EventStatusPending),
					string(events.EventStatusFailed),
				}).
			Order("CASE status WHEN 'pending' THEN 1 WHEN 'failed' THEN 2 END, created_at ASC").
			First(ctx)

		if err != nil {
			return err
		}

		if err := e.Processing(); err != nil {
			return err
		}
		if err := tx.Save(&e).Error; err != nil {
			return err
		}

		eventModel = e
		return nil
	})

	if err != nil {
		return 0, nil, err
	}

	return eventModel.ID, eventModel.Payload, nil
}

func (c *consumer) MarkEventAsPending(ctx context.Context, id uint) error {
	err := c.db.Transaction(func(tx *gorm.DB) error {
		eventModel, err := gorm.G[event](tx, clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
			Where("id = ?", id).
			First(ctx)

		if err != nil {
			return err
		}

		if err := eventModel.Pending(); err != nil {
			return err
		}

		return tx.Save(&eventModel).Error
	})

	return err
}

func (c *consumer) MarkEventAsProcessing(ctx context.Context, id uint) error {
	err := c.db.Transaction(func(tx *gorm.DB) error {
		eventModel, err := gorm.G[event](tx, clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
			Where("id = ?", id).
			First(ctx)

		if err != nil {
			return err
		}

		if err := eventModel.Processing(); err != nil {
			return err
		}

		return tx.Save(&eventModel).Error
	})

	return err
}

func (c *consumer) MarkEventAsProcessed(ctx context.Context, id uint) error {
	err := c.db.Transaction(func(tx *gorm.DB) error {
		eventModel, err := gorm.G[event](tx, clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
			Where("id = ?", id).
			First(ctx)

		if err != nil {
			return err
		}

		if err := eventModel.Processed(); err != nil {
			return err
		}

		return tx.Save(&eventModel).Error
	})

	return err
}

func (c *consumer) MarkEventAsFailed(ctx context.Context, id uint) error {
	err := c.db.Transaction(func(tx *gorm.DB) error {
		eventModel, err := gorm.G[event](tx, clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
			Where("id = ?", id).
			First(ctx)

		if err != nil {
			return err
		}

		if err := eventModel.Failed(); err != nil {
			return err
		}

		return tx.Save(&eventModel).Error
	})

	return err
}
