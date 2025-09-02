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
	grand := grand.Grand{}

	res := r.db.WithContext(ctx).Model(&grand.Grand{}).Where("uid = ?").First(&grand)
	if err := res.Error; err != nil {
	}
}
