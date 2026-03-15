package infra

import (
	"gorm.io/gorm"
	"octodome.com/eventbroker/infrastructure/model"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Event{})
	db.AutoMigrate(&model.Handler{})
}
