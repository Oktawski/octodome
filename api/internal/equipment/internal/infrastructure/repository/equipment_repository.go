package repo

import (
	"context"
	"errors"

	"octodome.com/api/internal/core"
	corecontext "octodome.com/api/internal/core/context"
	domain "octodome.com/api/internal/equipment/internal/domain/equipment"
	"octodome.com/api/internal/equipment/internal/infrastructure/model"
	collection "octodome.com/shared/collection"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type pgEquipmentRepository struct {
	db *gorm.DB
}

func NewPgEquipmentRepository(db *gorm.DB) *pgEquipmentRepository {
	return &pgEquipmentRepository{db: db}
}

func (r *pgEquipmentRepository) GetList(
	ctx context.Context,
	page int,
	pageSize int,
) ([]domain.Equipment, int64, error) {
	query := gorm.G[model.Equipment](r.db).Where("user_id = ?", ctx.Value(corecontext.UserIDKey))

	count, err := query.Count(ctx, "id")
	equipments, err := query.Scopes(core.PaginateG(page, pageSize)).Find(ctx)

	if err != nil {
		return nil, 0, err
	}

	equipmentList := collection.Map(
		equipments,
		func(e model.Equipment) domain.Equipment {
			return *e.ToDomain()
		},
	)

	return equipmentList, count, nil
}

func (r *pgEquipmentRepository) GetByID(
	ctx context.Context,
	id uint,
) (*domain.Equipment, error) {
	equipment, err := gorm.G[model.Equipment](r.db).
		Where("equipments.id = ?", id).
		Joins(clause.LeftJoin.Association("EquipmentType"), nil).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return equipment.ToDomain(), nil
}

func (r *pgEquipmentRepository) Create(
	ctx context.Context,
	e *domain.Equipment,
) error {
	equipment := model.EquipmentFromDomain(e)

	if err := gorm.G[model.Equipment](r.db).Create(ctx, equipment); err != nil {
		return err
	}

	return nil
}

func (r *pgEquipmentRepository) Update(
	ctx context.Context,
	e *domain.Equipment,
) error {
	equipment, err := gorm.G[model.Equipment](r.db).
		Where("id = ? AND user_id = ?", e.ID, e.UserID).
		First(ctx)
	if err != nil {
		return err
	}

	ct := equipment.Update(e.UserID, e)

	if ct.HasChanges {
		_, err = gorm.G[model.Equipment](r.db).Updates(ctx, equipment)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *pgEquipmentRepository) Delete(
	ctx context.Context,
	id uint,
) error {
	rowsAffected, err := gorm.G[model.Equipment](r.db).
		Where("id = ?", id).
		Delete(ctx)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("equipment not found or not owned by the user")
	}

	return nil
}

func (r *pgEquipmentRepository) ExistsByNameAndType(
	ctx context.Context,
	name string,
	equipmentTypeID uint,
) bool {
	var count int64

	if err := r.db.Model(&model.Equipment{}).
		Where("name = ?", name).
		Where("type_id = ?", equipmentTypeID).
		Where("user_id = ?", ctx.Value(corecontext.UserIDKey)).
		Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}

func (r *pgEquipmentRepository) IsOwnedByUser(
	ctx context.Context,
	equipmentID uint,
) bool {
	var count int64

	if err := r.db.Model(&model.Equipment{}).
		Where("id = ?", equipmentID).
		Where("user_id = ?", ctx.Value(corecontext.UserIDKey)).
		Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}
