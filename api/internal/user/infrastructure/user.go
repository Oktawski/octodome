package infra

import (
	"octodome.com/api/internal/user/domain"
	"octodome.com/shared/valuetype"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        valuetype.Email `gorm:"column:email;uniqueIndex;not null"`
	Username     string          `gorm:"column:username;"`
	PasswordHash string          `gorm:"column:password;not null"`
}

func (User) TableName() string {
	return "users"
}

func (e *User) ToDomain() *domain.User {
	return &domain.User{
		ID:           e.ID,
		Username:     e.Username,
		Email:        e.Email,
		PasswordHash: e.PasswordHash,
	}
}

func FromDomain(u *domain.User) *User {
	return &User{
		Model:        gorm.Model{ID: u.ID},
		Email:        u.Email,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}
}
