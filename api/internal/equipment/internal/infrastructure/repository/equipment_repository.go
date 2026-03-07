package repo

import (
	"context"

	authdom "octodome.com/api/internal/auth/domain"
	"octodome.com/api/internal/core"
	domain "octodome.com/api/internal/equipment/internal/domain/equipment"
	"octodome.com/api/internal/equipment/internal/infrastructure/model"
	collection "octodome.com/shared/collection"

	"gorm.io/gorm"
)

type pgEquipmentRepository struct {
	db *gorm.DB
}

func NewPgEquipmentRepository(db *gorm.DB) *pgEquipmentRepository {
	return &pgEquipmentRepository{db: db}
}

func (r *pgEquipmentRepository) GetList(
	userContext authdom.UserContext,
	page int,
	pageSize int,
) ([]domain.Equipment, int64, error) {

	var equipments []model.Equipment

	var count int64

	if dbError := r.db.
		Where("user_id = ?", userContext.ID).
		Scopes(core.Paginate(page, pageSize)).
		Find(&equipments).
		Count(&count).Error; dbError != nil {
		return nil, 0, dbError
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
	userContext authdom.UserContext,
	id uint,
) (*domain.Equipment, error) {

	var eq *model.Equipment

	// TODO: this doesn't include the equipment type 2026-03-07
	if dbError := r.db.
		Where("id = ? AND user_id = ?", id, userContext.ID).
		First(&eq).Error; dbError != nil {
		return nil, dbError
	}

	return eq.ToDomain(), nil
}

func (r *pgEquipmentRepository) Create(e *domain.Equipment) error {
	equipment := model.EquipmentFromDomain(e)

	if dbError := r.db.Create(&equipment).Error; dbError != nil {
		return dbError
	}

	return nil
}

func (r *pgEquipmentRepository) Update(
	userContext authdom.UserContext,
	ctx context.Context,
	e *domain.Equipment,
) error {
	equipment, err := gorm.G[model.Equipment](r.db).Where("id = ? AND user_id = ?", e.ID, e.UserID).First(ctx)
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

func (r *pgEquipmentRepository) Delete(id uint) error {
	if dbError := r.db.
		Model(&model.Equipment{}).
		Where("ID = ?", id).
		Delete(&model.Equipment{}).Error; dbError != nil {
		return dbError
	}
	return nil
}

func (r *pgEquipmentRepository) ExistsByNameAndType(
	userContext authdom.UserContext,
	name string,
	equipmentTypeID uint,
) bool {
	var count int64

	if err := r.db.Model(&model.Equipment{}).
		Where("name = ?", name).
		Where("type_id = ?", equipmentTypeID).
		Where("user_id = ?", userContext.ID).
		Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}

func (r *pgEquipmentRepository) IsOwnedByUser(
	userContext authdom.UserContext,
	equipmentID uint,
) bool {
	var count int64

	if err := r.db.Model(&model.Equipment{}).
		Where("id = ?", equipmentID).
		Where("user_id = ?", userContext.ID).
		Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}
