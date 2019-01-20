package config

import (
	"github.com/kelseyhightower/envconfig"

	"fmt"
)

type Config struct {
	DB DB
}

type DB struct {
	Username string
	Password string
	Instance string
	Schema   string
}

var (
	config *Config
)

func NewConfig() (*Config, error) {
	if config == nil {
		config = &Config{}
		if err := envconfig.Process("config", config); err != nil {
			return nil, fmt.Errorf("could not read env: %v",err)
		}
	}
	return config, nil
}
