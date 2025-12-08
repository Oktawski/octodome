package infra

import (
	"gorm.io/gorm"
	"octodome.com/eventbroker/domain"
	"octodome.com/shared/events"
)

type consumer struct {
	db *gorm.DB
}

func NewEventConsumer(db *gorm.DB) domain.Consumer {
	return &consumer{db: db}
}

func (c *consumer) GetEvent(eventType string) (uint, interface{}, error) {
	var event event
	err := c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Set("gorm:query_option", "FOR UPDATE SKIP LOCKED").
			Where("type = ? AND status IN ?",
				eventType,
				[]string{
					string(events.EventStatusPending),
					string(events.EventStatusFailed),
				}).
			Order("CASE status WHEN 'pending' THEN 1 WHEN 'failed' THEN 2 END, created_at ASC").
			First(&event).Error; err != nil {
			return err
		}

		if err := event.Processing(); err != nil {
			return err
		}

		return tx.Save(&event).Error
	})

	if err != nil {
		return 0, nil, err
	}

	return event.ID, event.Payload, nil
}

func (c *consumer) MarkEventAsPending(id uint) error {
	var event event

	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Set("gorm:query_option", "FOR UPDATE").
			Where("id = ?", id).
			First(&event).Error; err != nil {
			return err
		}

		if err := event.Pending(); err != nil {
			return err
		}

		return tx.Save(&event).Error
	})
}

func (c *consumer) MarkEventAsProcessing(id uint) error {
	var event event

	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Set("gorm:query_option", "FOR UPDATE").
			Where("id = ?", id).
			First(&event).Error; err != nil {
			return err
		}

		if err := event.Processing(); err != nil {
			return err
		}

		return tx.Save(&event).Error
	})
}

func (c *consumer) MarkEventAsProcessed(id uint) error {
	var event event

	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Set("gorm:query_option", "FOR UPDATE").
			Where("id = ?", id).
			First(&event).Error; err != nil {
			return err
		}

		if err := event.Processed(); err != nil {
			return err
		}

		return tx.Save(&event).Error
	})
}

func (c *consumer) MarkEventAsFailed(id uint) error {
	var event event

	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Set("gorm:query_option", "FOR UPDATE").
			Where("id = ?", id).
			First(&event).Error; err != nil {
			return err
		}

		if err := event.Failed(); err != nil {
			return err
		}

		return tx.Save(&event).Error
	})
}
