package config

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr            string `envconfig:"ADDR" default:"localhost:50051"`
	RouteUserRole   string `envconfig:"ROUTE_USER_ROLE" default:"admin"`
	RouteGrandRole  string `envconfig:"ROUTE_GRAND_ROLE" default:"admin"`
	RouteDataRole   string `envconfig:"ROUTE_DATA_ROLE" default:"admin"`
	DataChiferKey   string `envconfig:"DATA_CHIFER_KEY" default:"secret"`
	RedisUrl        string `encconfig:"REDIS_URL" default:"localhost:9000"`
	ElasticSearhUrl string `encconfig:"ELASTICSEARCH_URL" default:"http://localhost:9200"`
	DataIndexName   string `envconfig:"DATA_INDEX_NAME" default:"data"`
	DSN             string
}

func NewConfig() (*Config, error) {
	cfg := Config{}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	cfg.DSN = dsn

	return &cfg, nil
}
