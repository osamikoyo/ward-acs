package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/osamikoyo/ward/entity/grand"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *Repository) CreateGrand(ctx context.Context, grand *grand.Grand) error {
	if grand == nil {
		return ErrNilInput
	}

	if err := r.db.WithContext(ctx).Create(grand).Error; err != nil {
		r.logger.Error("failed create grand",
			zap.Any("grand", grand),
			zap.Error(err))

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ErrAlreadyExist
		}

		return ErrUnknown
	}

	r.logger.Info("grand created successfully", zap.Any("grand", grand))

	return nil
}

func (r *Repository) UpdateGrand(ctx context.Context, uid uuid.UUID, field string, value interface{}) error {
	res := r.db.WithContext(ctx).Model(&grand.Grand{}).Where("uid = ?", uid).Update(field, value)

	if res.RowsAffected == 0 {
		return ErrNotFound
	}
	if err := res.Error; err != nil {
		r.logger.Error("failed update grand",
			zap.String("uid", uid.String()),
			zap.String("field", field),
			zap.Error(err))

		return ErrUnknown
	}

	r.logger.Info("grand updated successfully", zap.String("uid", uid.String()))

	return nil
}

func (r *Repository) GetGrandByUUID(ctx context.Context, uid uuid.UUID) (*grand.Grand, error) {
	fethedgrand := grand.Grand{}

	res := r.db.WithContext(ctx).Model(&grand.Grand{}).Where("uid = ?").First(&fethedgrand)
	if err := res.Error; err != nil {
		r.logger.Error("failed fetch grand by uuid",
			zap.String("uuid", uid.String()),
			zap.Error(err))

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, ErrAlreadyExist
		}

		return nil, ErrUnknown
	}

	r.logger.Info("successfully fetched grand",
		zap.String("uid", uid.String()))

	return &fethedgrand, nil
}

func (r *Repository) DeleteGrand(ctx context.Context, uid uuid.UUID) error {
	if err := r.db.Where("uid = ?", uid).Delete(&grand.Grand{}).Error; err != nil {
		r.logger.Error("faile delete grand",
			zap.String("uid", uid.String()),
			zap.Error(err))

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}

		return ErrUnknown
	}

	r.logger.Error("grand deleted successfully", zap.String("uid", uid.String()))

	return nil
}

func (r *Repository) ListGrands(ctx context.Context) ([]grand.Grand, error) {
	grands := []grand.Grand{}

	res := r.db.WithContext(ctx).Find(&grands)
	if err := res.Error; err != nil {
		r.logger.Error("failed list grands", zap.Error(err))

		return nil, ErrUnknown
	}

	r.logger.Info("grands successfully listed")

	return grands, nil
}
