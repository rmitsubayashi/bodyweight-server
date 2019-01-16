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
}

var (
	config *Config
)

func NewConfig() *Config {
	if config == nil {
		config = &Config{}
		err := envconfig.Process("config", config)
		if err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return config
}
