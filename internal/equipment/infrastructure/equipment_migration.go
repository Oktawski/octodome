package eqinfra

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&equipment{})
	db.AutoMigrate(&equipmentType{})
}
