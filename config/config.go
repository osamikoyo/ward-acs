package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Addr           string `envconfig:"ADDR" default:"localhost:50051"`
	RouteUserRole  string `envconfig:"ROUTE_USER_ROLE" default:"admin"`
	RouteGrandRole string `envconfig:"ROUTE_GRAND_ROLE" default:"admin"`
	RouteDataRole  string `envconfig:"ROUTE_DATA_ROLE" default:"admin"`
	DSN            string
}

func NewConfig() (*Config, error) {
	cfg := Config{}

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
