package infra

import (
	"gorm.io/gorm"
	"octodome.com/eventbroker"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&eventbroker.Event{})
}
