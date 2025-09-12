package searchbase

import (
	"bytes"
	"context"

	"github.com/elastic/go-elasticsearch/v8"
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

	return nil
}
