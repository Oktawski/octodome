package infra

import (
	infra "octodome/internal/user/infrastructure"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&infra.User{})
}
