package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"octodome.com/api/internal/auth/internal/infrastructure/model"
)

type pgMagicCode struct {
	db *gorm.DB
}

func NewPgMagicCode(db *gorm.DB) *pgMagicCode {
	return &pgMagicCode{db: db}
}

func (r *pgMagicCode) Create(ctx context.Context, code string, email string) error {
	return gorm.G[model.MagicCode](r.db).Create(ctx, model.NewMagicCode(code, email))
}

func (r *pgMagicCode) GetByEmailAndCode(ctx context.Context, email string, code string) (string, error) {
	magicCode, err := gorm.G[model.MagicCode](r.db).Where("email = ? AND code = ?", email, code).First(ctx)
	if err != nil {
		return "", err
	}

	return magicCode.Code, nil
}

func (r *pgMagicCode) DeleteByEmail(ctx context.Context, email string) error {
	rowsAffected, err := gorm.G[model.MagicCode](r.db).Where("email = ?", email).Delete(ctx)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("magic code not found")
	}

	return nil
}
