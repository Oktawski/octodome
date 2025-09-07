package infra

import (
	"octodome/internal/equipment/infrastructure/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Equipment{})
	db.AutoMigrate(&model.EquipmentType{})
}
