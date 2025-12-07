package migration

import (
	"octodome.com/api/internal/auth/internal/infrastructure/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Role{}, &model.UserRole{})
}
