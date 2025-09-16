package model

type Role struct {
	Name string `gorm:"uniqueIndex;not null;primaryKey"`
}

type UserRole struct {
	UserID uint   `gorm:"uniqueIndex;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID"`
	Role   string `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Role;references:Name"`
}
