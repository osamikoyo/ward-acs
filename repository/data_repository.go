package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/osamikoyo/ward/entity/data"
	"github.com/osamikoyo/ward/entity/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *Repository) CreateData(ctx context.Context, data *data.Data) error {
	if data == nil {
		return ErrNilInput
	}

	if err := r.db.WithContext(ctx).Create(data).Error; err != nil {
		r.logger.Error("failed create data",
			zap.Any("data", data),
			zap.Error(err))

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ErrAlreadyExist
		}

		return ErrUnknown
	}

	r.logger.Info("data created successfully", zap.Any("data", data))

	return nil
}

func (r *Repository) GetData(ctx context.Context, uid uuid.UUID) (*data.Data, error) {
	data := data.Data{}

	if err := r.db.WithContext(ctx).Where("uid = ?", uid).First(&data).Error; err != nil {
		r.logger.Error("failed fetch data",
			zap.String("uid", uid.String()),
			zap.Error(err))

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, ErrUnknown
	}

	r.logger.Info("data fetched successfully", zap.Any("data", data))

	return &data, nil
}

func (r *Repository) UpdateData(ctx context.Context, uid uuid.UUID, field string, value interface{}) error {
	res := r.db.WithContext(ctx).Model(&user.User{}).Where("uid = ?", uid).Update(field, value)
	if res.RowsAffected == 0 {
		return ErrNotFound
	}
	if err := res.Error; err != nil {
		r.logger.Error("failed update data",
			zap.String("uid", uid.String()),
			zap.String("field", field),
			zap.Error(err))

		return ErrUnknown
	}

	r.logger.Info("data updated successfully",
		zap.String("uid", uid.String()),
		zap.String("field", field))

	return nil
}

func (r *Repository) ListData(ctx context.Context, grandUID uuid.UUID) ([]data.Data, error) {

}

func (r *Repository) DeleteData(ctx context.Context, uid uuid.UUID) error {
	res := r.db.WithContext(ctx).Where("uid = ?", uid).Delete(&data.Data{})
	if res.RowsAffected == 0 {
		return ErrNotFound
	}
	if err := res.Error; err != nil {
		r.logger.Error("failed delete data",
			zap.String("uid", uid.String()),
			zap.Error(err))

		return ErrUnknown
	}

	r.logger.Info("data deleted successfully", zap.String("uid", uid.String()))

	return nil
}
