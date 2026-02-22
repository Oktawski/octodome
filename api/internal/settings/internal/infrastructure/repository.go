package infra

import (
	"context"
	"encoding/json"
	"time"

	"gorm.io/gorm"
	"octodome.com/api/internal/settings/internal/domain"
)

type pgSettingRepository struct {
	db *gorm.DB
}

func NewPgSettingRepository(db *gorm.DB) *pgSettingRepository {
	return &pgSettingRepository{db: db}
}

func (r *pgSettingRepository) Get(ctx context.Context, name string) (*domain.Setting, error) {
	setting, err := gorm.G[Setting](r.db).
		Where("name = ?", name).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return setting.ToDomain(), nil
}

func (r *pgSettingRepository) Set(ctx context.Context, setting domain.Setting) error {
	rowsAffected, err := gorm.G[Setting](r.db).
		Where("name = ?", setting.Name).
		Updates(ctx, Setting{
			Value:     json.RawMessage(setting.Value),
			UpdatedAt: time.Now(),
		})
	if err != nil {
		return err
	}

	if rowsAffected > 0 {
		return nil
	}

	settingModel := NewSetting(setting)
	return gorm.G[Setting](r.db).Create(ctx, settingModel)
}
