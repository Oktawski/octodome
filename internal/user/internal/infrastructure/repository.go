package infra

import (
	"octodome/internal/user/domain"

	"gorm.io/gorm"
)

type pgUserRepository struct {
	db *gorm.DB
}

func NewPgUserRepository(db *gorm.DB) *pgUserRepository {
	return &pgUserRepository{db: db}
}

func (r *pgUserRepository) GetByID(id uint) (*domain.User, error) {
	var u User

	result := r.db.First(&u, id)

	return u.ToDomain(), result.Error
}

func (r *pgUserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var userModel *User

	dbError := r.db.Where("username = ?", username).First(&userModel).Error
	if dbError != nil {
		return nil, dbError
	}

	return userModel.ToDomain(), nil
}

func (r *pgUserRepository) Create(user *domain.User) error {
	gormUser := fromDomain(user)

	if err := r.db.Create(gormUser).Error; err != nil {
		return err
	}

	return nil
}
