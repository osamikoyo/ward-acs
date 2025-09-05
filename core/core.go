package core

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/osamikoyo/ward/config"
	"github.com/osamikoyo/ward/entity/data"
	"github.com/osamikoyo/ward/entity/grand"
	"github.com/osamikoyo/ward/entity/user"
	"github.com/osamikoyo/ward/logger"
	"github.com/osamikoyo/ward/repository"
)

var ErrPermissionDenied = errors.New("permission denied")

const TokenLength = 40

type WardCore struct {
	repository *repository.Repository
	logger     *logger.Logger

	timeout time.Duration
	cfg     *config.Config
}

func NewWardCore(repository *repository.Repository, logger *logger.Logger, cfg *config.Config, timeout time.Duration) *WardCore {
	return &WardCore{
		repository: repository,
		logger:     logger,
		timeout:    timeout,
		cfg:        cfg,
	}
}

func (w *WardCore) context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), w.timeout)
}

func (w *WardCore) generateToken() (string, error) {
	b := make([]byte, TokenLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:TokenLength], nil
}

func (w *WardCore) RegisterUser(token string, grandUID uuid.UUID) (string, error) {
	ctx, cancel := w.context()
	defer cancel()

	requser, err := w.repository.GetUserByToken(ctx, token)
	if err != nil {
		return "", err
	}

	if requser.Grand.Name != w.cfg.RouteUserRole {
		return "", ErrPermissionDenied
	}

	tokenOut, err := w.generateToken()
	if err != nil {
		return "", fmt.Errorf("failed generate token: %v", err)
	}

	user := user.NewUser(token, grandUID)

	if err := w.repository.CreateUser(ctx, user); err != nil {
		return "", err
	}

	return tokenOut, nil
}

func (w *WardCore) ChangeUserGrand(token string, userUID uuid.UUID, grandUID uuid.UUID) error {
	ctx, cancel := w.context()
	defer cancel()

	reqUser, err := w.repository.GetUserByToken(ctx, token)
	if err != nil {
		return err
	}

	if reqUser.Grand.Name != w.cfg.RouteUserRole {
		return ErrPermissionDenied
	}

	if err = w.repository.UpdateUser(ctx, userUID, "grand_uid", grandUID); err != nil {
		return err
	}

	return nil
}

func (w *WardCore) GetData(token string, dataUID uuid.UUID) (*data.Data, error) {
	ctx, cancel := w.context()
	defer cancel()

	user, err := w.repository.GetUserByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	data, err := w.repository.GetData(ctx, dataUID)
	if err != nil {
		return nil, err
	}

	if user.Grand.Level < data.Grand.Level {
		return nil, ErrPermissionDenied
	}

	return data, nil
}

func (w *WardCore) ChangeDataGrand(token string, dataUID, grandUID uuid.UUID) error {
	ctx, cancel := w.context()
	defer cancel()

	user, err := w.repository.GetUserByToken(ctx, token)
	if err != nil {
		return err
	}

	if user.Grand.Name != w.cfg.RouteGrandRole {
		return ErrPermissionDenied
	}

	if err := w.repository.UpdateData(ctx, dataUID, "grand_uid", grandUID); err != nil {
		return err
	}

	return nil
}

func (w *WardCore) CreateGrand(token, name string, level int) error {
	ctx, cancel := w.context()
	defer cancel()

	user, err := w.repository.GetUserByToken(ctx, token)
	if err != nil {
		return err
	}

	if user.Grand.Name != w.cfg.RouteGrandRole {
		return ErrPermissionDenied
	}

	grand := grand.NewGrand(name, level)

	return w.repository.CreateGrand(ctx, grand)
}
