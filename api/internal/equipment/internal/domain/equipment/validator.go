package domain

import (
	"context"
)

type Validator interface {
	CanBeCreated(
		ctx context.Context,
		name string,
		equipmentTypeID uint,
	) bool

	CanBeModified(
		ctx context.Context,
		equipmentID uint,
	) bool
}

type validator struct {
	repo Repository
}

func NewValidator(repo Repository) *validator {
	return &validator{repo: repo}
}

func (v validator) CanBeCreated(
	ctx context.Context,
	name string,
	equipmentTypeID uint,
) bool {
	return !v.repo.ExistsByNameAndType(ctx, name, equipmentTypeID)
}

func (v validator) CanBeModified(
	ctx context.Context,
	equipmentID uint,
) bool {
	return v.repo.IsOwnedByUser(ctx, equipmentID)
}
