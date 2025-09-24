package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/osamikoyo/ward/entity/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *Repository) CreateUser(ctx context.Context, usr *user.User) error {
	if usr == nil {
		return ErrNilInput
	}

	if err := r.db.WithContext(ctx).Create(usr).Error; err != nil {
		r.logger.Error("failed create user",
			zap.Any("user", usr),
			zap.Error(err))

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ErrAlreadyExist
		}

		return ErrUnknown
	}

	r.logger.Info("user created successfully", zap.Any("user", usr))

	return nil
}

func (r *Repository) ListUsers(ctx context.Context) ([]user.User, error) {

	users := []user.User{}

	res := r.db.WithContext(ctx).Find(&users)
	if err := res.Error; err != nil {
		r.logger.Error("failed list users", zap.Error(err))

		return nil, ErrUnknown
	}

	r.logger.Info("users listed successfully")

	return users, nil
}

func (r *Repository) ReadUsers(ctx context.Context, filter map[string]interface{}) ([]user.User, error) {
	users := []user.User{}

	res := r.db.WithContext(ctx).Where(filter).Find(&users)
	if err := res.Error; err != nil {
		r.logger.Error("failed read users",
			zap.Any("filter", filter),
			zap.Error(err))

		if res.RowsAffected == 0 {
			return nil, ErrNotFound
		}

		return nil, ErrUnknown
	}

	r.logger.Info("users fetched successfully", zap.Any("users", users))

	return users, nil
}

func (r *Repository) GetUserByToken(ctx context.Context, token string) (*user.User, error) {
	if len(token) == 0 {
		return nil, ErrNilInput
	}
	user := user.User{}

	if err := r.db.WithContext(ctx).Where("token = ?", token).First(user).Error; err != nil {
		r.logger.Error("failed fetch user",
			zap.String("token", token),
			zap.Error(err))

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}

		return nil, ErrUnknown
	}

	r.logger.Info("user fetched successfully",
		zap.Any("user", user),
	)

	return &user, nil
}

func (r *Repository) UpdateUser(ctx context.Context, uid uuid.UUID, field string, value interface{}) error {
	res := r.db.WithContext(ctx).Model(&user.User{}).Where("uid = ?", uid).Update(field, value)
	if err := res.Error; err != nil {
		r.logger.Error("failed update user",
			zap.String("field", field),
			zap.String("uid", uid.String()),
			zap.Error(err))

		if res.RowsAffected == 0 {
			return ErrNotFound
		}

		r.logger.Info("failed update user",
			zap.String("uid", uid.String()),
			zap.String("field", field))

		return ErrUnknown
	}

	r.logger.Info("user updated successfully",
		zap.String("uid", uid.String()),
		zap.String("field", field))

	return nil
}

func (r *Repository) DeleteUser(ctx context.Context, uid uuid.UUID) error {
	res := r.db.Where("uid = ?", uid).Delete(&user.User{})
	if err := res.Error; err != nil {
		r.logger.Error("failed delete user ",
			zap.String("uid", uid.String()),
			zap.Error(err))

		if res.RowsAffected == 0 {
			return ErrNotFound
		}

		return ErrUnknown
	}

	r.logger.Info("user deleted successfully", zap.String("uid", uid.String()))

	return nil
}
