package infra

import (
	"encoding/json"
	"time"

	"octodome.com/api/internal/settings/internal/domain"
)

type Setting struct {
	ID        uint            `gorm:"primaryKey"`
	Name      string          `gorm:"uniqueIndex;not null;primaryKey"`
	Value     json.RawMessage `gorm:"type:jsonb"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime"`
}

func (Setting) TableName() string {
	return "settings"
}

func NewSetting(setting domain.Setting) *Setting {
	return &Setting{
		Name:      setting.Name,
		Value:     json.RawMessage(setting.Value),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (s *Setting) ToDomain() *domain.Setting {
	return &domain.Setting{
		Name:  s.Name,
		Value: string(s.Value),
	}
}
