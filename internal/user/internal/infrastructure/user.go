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

func (g *User) ToDomain() *domain.User {
	return &domain.User{
		ID:           g.ID,
		Username:     g.Username,
		Email:        g.Email,
		PasswordHash: g.PasswordHash,
	}
}

func fromDomain(u *domain.User) *User {
	return &User{
		Model:        gorm.Model{ID: u.ID},
		Email:        u.Email,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}
}
