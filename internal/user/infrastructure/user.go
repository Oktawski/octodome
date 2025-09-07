package infra

import (
	domain "octodome/internal/user/domain"

	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	Email        string `gorm:"column:email"`
	Username     string `gorm:"column:username;uniqueIndex"`
	PasswordHash string `gorm:"column:password"`
}

func (user) TableName() string {
	return "users"
}

func (g *user) toDomain() *domain.User {
	return &domain.User{
		ID:           g.ID,
		Username:     g.Username,
		Email:        g.Email,
		PasswordHash: g.PasswordHash,
	}
}

func fromDomain(u *domain.User) *user {
	return &user{
		Model:        gorm.Model{ID: u.ID},
		Email:        u.Email,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}
}
