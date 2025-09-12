package searchbase

import (
	"bytes"
	"context"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
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
	req := esapi.IndexRequest{
		Index:   index,
		Body:    bytes.NewReader(jsonReq),
		Refresh: "true",
	}
	res, err := req.Do(ctx, s.client)

	if err != nil || res.IsError() {
		s.logger.Error("failed add value to index",
			zap.String("index", index),
			zap.Error(err))

		return err
	}
	defer res.Body.Close()

	s.logger.Info("value was added to index successfully", zap.String("index", index))

	return nil
}
