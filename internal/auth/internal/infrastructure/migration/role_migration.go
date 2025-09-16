package migration

import (
	"octodome/internal/auth/internal/infrastructure/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Role{}, &model.UserRole{})
}
