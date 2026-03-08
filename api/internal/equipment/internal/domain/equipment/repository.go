package domain

import (
	"context"
)

type Repository interface {
	GetList(ctx context.Context, page int, pageSize int) ([]Equipment, int64, error)
	GetByID(ctx context.Context, id uint) (*Equipment, error)
	Create(ctx context.Context, equipment *Equipment) error
	Update(ctx context.Context, equipment *Equipment) error
	Delete(ctx context.Context, id uint) error
	ExistsByNameAndType(ctx context.Context, name string, equipmentTypeID uint) bool
	IsOwnedByUser(ctx context.Context, equipmentID uint) bool
}
