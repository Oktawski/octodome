package model

import (
	"gorm.io/gorm"
)

type Handler struct {
	gorm.Model
	Name      string `json:"name"       gorm:"uniqueIndex:idx_handler_name_event_type"`
	EventType string `json:"event_type" gorm:"uniqueIndex:idx_handler_name_event_type"`
	URL       string `json:"url"`
}

func (h *Handler) TableName() string {
	return "handlers"
}
