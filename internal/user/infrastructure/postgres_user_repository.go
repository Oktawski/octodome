package userinfra

import (
	userdom "octodome/internal/user/domain"

	"gorm.io/gorm"
)

type pgUserRepository struct {
	db *gorm.DB
}

func NewPgUserRepository(db *gorm.DB) *pgUserRepository {
	return &pgUserRepository{db: db}
}

func (r *pgUserRepository) GetByID(id uint) (*userdom.User, error) {
	var gormUser gormUser

	result := r.db.First(&gormUser, id)

	return gormUser.toDomain(), result.Error
}

func (r *pgUserRepository) GetUserByUsername(username string) (user *userdom.User, err error) {
	var userModel *gormUser

	dbError := r.db.Where("username = ?", username).First(&userModel).Error
	if dbError != nil {
		return nil, dbError
	}

	return userModel.toDomain(), nil
}

func (r *pgUserRepository) Create(user *userdom.User) error {
	gormUser := fromDomain(user)

	if err := r.db.Create(gormUser).Error; err != nil {
		return err
	}

	return nil
}
