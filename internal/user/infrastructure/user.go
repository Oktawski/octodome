package infra

import (
	"octodome/internal/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"column:email"`
	Username     string `gorm:"column:username;uniqueIndex"`
	PasswordHash string `gorm:"column:password"`
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
