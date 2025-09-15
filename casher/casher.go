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

func (c *Casher) GetFromCash(ctx context.Context, key string) (string, error) {
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		c.logger.Error("failed fetch from cash",
			zap.String("key", key),
			zap.Error(err))

		return "", err
	}

	c.logger.Info("value was successfully fetched", zap.String("key", key))

	return value, nil
}

func (c *Casher) DeleteFromCash(ctx context.Context, key string) error {
	_, err := c.client.Del(ctx, key).Result()
	if err != nil {
		c.logger.Error("failed delete from cash",
			zap.String("key", key),
			zap.Error(err))

		return err
	}

	c.logger.Info("value was successfully deleted from cash", zap.String("key", key))

	return nil
}
