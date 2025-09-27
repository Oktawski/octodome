package model

import infra "octodome/internal/user/infrastructure"

type Role struct {
	Name string `gorm:"uniqueIndex;not null;primaryKey"`
}

type UserRole struct {
	UserID uint
	User   infra.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID"`

	RoleID string `gorm:"not null;index"`
	Role   Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RoleID;references:Name"`
}
