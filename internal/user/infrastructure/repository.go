package infra

import (
	authdom "octodome/internal/auth/domain"
	infra "octodome/internal/user/internal/infrastructure"

	"gorm.io/gorm"
)

type pgUserRepository struct {
	db *gorm.DB
}

func NewPgUserRepository(db *gorm.DB) *pgUserRepository {
	return &pgUserRepository{db: db}
}

func (r *pgUserRepository) GetUserAuthDTO(username string) (*authdom.UserAuthDTO, error) {
	var user authdom.UserAuthDTO

	dbError := r.db.
		Model(infra.User{}).
		Select("id, username, password").
		Where("username = ?", username).
		Take(&user).
		Error
	if dbError != nil {
		return nil, dbError
	}

	return &user, nil
}
