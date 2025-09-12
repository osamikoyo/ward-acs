package casher

import "github.com/redis/go-redis/v9"

type Casher struct {
	client *redis.Client
}
