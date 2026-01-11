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
	user, err := gorm.G[User](r.db).Where("username = ?", username).First(ctx)
	if err != nil {
		return nil, err
	}

	return &authdom.UserAuthDTO{
		ID:       user.ID,
		Username: user.Username,
		Password: user.PasswordHash,
	}, nil
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
