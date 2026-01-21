package infra

import (
	"context"

	"octodome.com/api/internal/user/domain"
	infra "octodome.com/api/internal/user/infrastructure"

	"gorm.io/gorm"
)

type pgUserRepository struct {
	db *gorm.DB
}

func NewPgUserRepository(db *gorm.DB) *pgUserRepository {
	return &pgUserRepository{db: db}
}

func (r *pgUserRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	var u infra.User

	result := r.db.First(&u, id)

	return u.ToDomain(), result.Error
}

func (r *pgUserRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	user, err := gorm.G[infra.User](r.db).Where("username = ?", username).First(ctx)
	if err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}

func (r *pgUserRepository) Create(ctx context.Context, user *domain.User) (uint, error) {
	userModel := infra.FromDomain(user)

	if err := gorm.G[infra.User](r.db).Create(ctx, userModel); err != nil {
		return 0, err
	}

	return userModel.ID, nil
}

func (r *pgUserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	count, err := gorm.G[infra.User](r.db).
		Where("email = ?", email).
		Count(ctx, "ID")
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
