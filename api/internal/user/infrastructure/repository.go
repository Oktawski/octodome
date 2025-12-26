package infra

import (
	"context"
	"errors"

	authdom "octodome.com/api/internal/auth/domain"

	"gorm.io/gorm"
)

type pgUserRepository struct {
	db *gorm.DB
}

func NewPgUserRepository(db *gorm.DB) *pgUserRepository {
	return &pgUserRepository{db: db}
}

func (r *pgUserRepository) GetUserAuthDTO(ctx context.Context, username string) (*authdom.UserAuthDTO, error) {
	var user authdom.UserAuthDTO

	dbError := r.db.
		Model(User{}).
		Select("id, username, password").
		Where("username = ?", username).
		Take(&user).
		Error
	if dbError != nil {
		return nil, dbError
	}

	return &user, nil
}

func (r *pgUserRepository) ExistsByEmailOrUsername(
	ctx context.Context,
	email string,
	username string,
) (bool, error) {
	count, err := gorm.G[User](r.db).Where("email = ? or username = ?", email, username).Count(ctx, "ID")
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, errors.New("email or username already exists")
	}

	return false, nil
}
