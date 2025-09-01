package repository

import (
	"errors"

	"github.com/osamikoyo/ward/logger"
	"gorm.io/gorm"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrAlreadyExist = errors.New("already exist")
	ErrUnknown      = errors.New("repository error")
	ErrNilInput = errors.New("nil input")
)

type Repository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func NewRepository(db *gorm.DB, logger *logger.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}
