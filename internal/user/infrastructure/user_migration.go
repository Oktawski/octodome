package userinfra

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&gormUser{})
}
