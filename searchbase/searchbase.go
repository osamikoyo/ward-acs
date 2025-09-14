package searchbase

import (
	"bytes"
	"context"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/google/uuid"
	"github.com/osamikoyo/ward/logger"
	"go.uber.org/zap"
)

type SearchBase struct {
	client *elasticsearch.Client
	logger *logger.Logger
}

func NewSearchBase(client *elasticsearch.Client, logger *logger.Logger) *SearchBase {
	return &SearchBase{
		client: client,
		logger: logger,
	}
}

func (s *SearchBase) AddToSearchBase(ctx context.Context, index string, jsonReq []byte) error {
	res, err := s.client.Index(
		index,
		bytes.NewReader(jsonReq),
		s.client.Index.WithContext(ctx),
	)
	if err != nil {
		s.logger.Error("failed add to search base",
			zap.String("index", index),
			zap.Error(err))

		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		s.logger.Error("result with error", zap.String("res", res.String()))
	}

	s.logger.Info("value was successfully added to search base")

	return nil
}

func (s *SearchBase) DeleteFromSearchBase(ctx context.Context, index string, id uuid.UUID) error {
	res, err := s.client.Delete(
		index,
		id.String(),
		s.client.Delete.WithRefresh("true"),
		s.client.Delete.WithContext(ctx),
	)
	if err != nil {
		s.logger.Error("failed delete value from search base",
			zap.String("index", index),
			zap.String("id", id.String()),
			zap.Error(err))

		return err
	}

	if res.IsError() {
		s.logger.Debug("result is err", zap.String("res", res.String()))
	}

	s.logger.Info("value was deleted from search base successfully")

	return nil
}
