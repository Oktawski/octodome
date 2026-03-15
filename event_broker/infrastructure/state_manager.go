package infra

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"octodome.com/eventbroker/domain"
	"octodome.com/eventbroker/infrastructure/model"
)

type stateManager struct {
	db *gorm.DB
}

func NewStateManager(db *gorm.DB) domain.StateManager {
	return &stateManager{db: db}
}

func (c *stateManager) MarkEventAsPending(ctx context.Context, id uint) error {
	err := c.db.Transaction(func(tx *gorm.DB) error {
		eventModel, err := gorm.G[model.Event](tx, clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
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

func (c *stateManager) MarkEventAsProcessing(ctx context.Context, id uint) error {
	err := c.db.Transaction(func(tx *gorm.DB) error {
		eventModel, err := gorm.G[model.Event](tx, clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
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

func (c *stateManager) MarkEventAsProcessed(ctx context.Context, id uint) error {
	err := c.db.Transaction(func(tx *gorm.DB) error {
		eventModel, err := gorm.G[model.Event](tx, clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
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

func (c *stateManager) MarkEventAsFailed(ctx context.Context, id uint) error {
	err := c.db.Transaction(func(tx *gorm.DB) error {
		eventModel, err := gorm.G[model.Event](tx, clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
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
