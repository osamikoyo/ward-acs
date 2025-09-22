package core

import (
	"context"

	"github.com/google/uuid"
	"github.com/osamikoyo/ward/entity/data"
	"github.com/osamikoyo/ward/entity/grand"
	"github.com/osamikoyo/ward/entity/user"
)

type (
	Repository interface {
		CreateUser(ctx context.Context, user *user.User) error
		GetUserByToken(ctx context.Context, token string) (*user.User, error)
		UpdateUser(ctx context.Context, uid uuid.UUID, field string, value interface{}) error
		DeleteUser(ctx context.Context, uid uuid.UUID) error

		CreateData(ctx context.Context, data *data.Data) error
		GetData(ctx context.Context, uid uuid.UUID) (*data.Data, error)
		UpdateData(ctx context.Context, uid uuid.UUID, field string, value interface{}) error
		ListData(ctx context.Context, level int) ([]data.Data, error)
		DeleteData(ctx context.Context, uid uuid.UUID) error

		CreateGrand(ctx context.Context, grand *grand.Grand) error
		UpdateGrand(ctx context.Context, uid uuid.UUID, field string, value interface{}) error
		GetGrandByUUID(ctx context.Context, uid uuid.UUID) (*grand.Grand, error)
		DeleteGrand(ctx context.Context, uid uuid.UUID) error
		ListGrands(ctx context.Context) ([]grand.Grand, error)
	}

	SearchBase interface {
		AddToSearchBase(ctx context.Context, index string, jsonReq []byte) error
		Search(ctx context.Context, index string, keywords []string) ([]byte, error)
		DeleteFromSearchBase(ctx context.Context, index string, id uuid.UUID) error
	}
)
