package searchbase

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

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

func (s *SearchBase) Search(ctx context.Context, index string, keywords []string) ([]byte, error) {
	keywordsInMap := make([]map[string]interface{}, len(keywords))
	for i, key := range keywords {
		keywordsInMap[i] = map[string]interface{}{
			"content": key,
		}
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": keywordsInMap,
			},
		},
	}

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		s.logger.Error("failed encode query",
			zap.Any("query", query),
			zap.Error(err))

		return nil, err
	}

	res, err := s.client.Search(
		s.client.Search.WithContext(ctx),
		s.client.Search.WithIndex(index),
		s.client.Search.WithBody(&buf),
		s.client.Search.WithPretty(),
	)
	if err != nil {
		s.logger.Error("failed send search query",
			zap.Any("query", query),
			zap.Error(err))

		return nil, err
	}

	defer res.Body.Close()

	if res.IsError() {
		s.logger.Warn("result is error",
			zap.String("res", res.String()))
	}

	result, err := io.ReadAll(res.Body)
	if err != nil {
		s.logger.Error("failed read response body",
			zap.Error(err))
	}

	return result, nil
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
