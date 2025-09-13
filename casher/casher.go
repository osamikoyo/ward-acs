package casher

import (
	"context"
	"time"

	"github.com/osamikoyo/ward/logger"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Casher struct {
	client *redis.Client
	logger *logger.Logger
}

var (
	ExpTime = time.Hour
)

func NewCasher(client *redis.Client, logger *logger.Logger) *Casher {
	return &Casher{
		client: client,
		logger: logger,
	}
}

func (c *Casher) AddToCash(ctx context.Context, key string, value string) error {
	_, err := c.client.Set(ctx, key, value, ExpTime).Result()
	if err != nil {
		c.logger.Error("failed add to cash",
			zap.String("key", key),
			zap.String("value", value),
			zap.Error(err))

		return err
	}

	c.logger.Info("value was added to cash successfully", zap.String("value", value))

	return nil
}
