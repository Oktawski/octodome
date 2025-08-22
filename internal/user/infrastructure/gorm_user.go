package userinfra

import (
	userdom "octodome/internal/user/domain"

	"gorm.io/gorm"
)

type gormUser struct {
	gorm.Model
	Email        string `gorm:"column:email"`
	Username     string `gorm:"column:username;uniqueIndex"`
	PasswordHash string `gorm:"column:password"`
}

func (gormUser) TableName() string {
	return "users"
}

func (g *gormUser) toDomain() *userdom.User {
	return &userdom.User{
		ID:           g.ID,
		Username:     g.Username,
		Email:        g.Email,
		PasswordHash: g.PasswordHash,
	}
}

func fromDomain(u *userdom.User) *gormUser {
	return &gormUser{
		Model: gorm.Model{
			ID: u.ID,
		},
		Email:        u.Email,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}
}
