package infra

import (
	infra "octodome.com/api/internal/user/infrastructure"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&infra.User{})
}
