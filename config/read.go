package config

import (
	"context"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
)

func ReadConfig(ctx context.Context) (*Config, error) {
	config.WithOptions(config.ParseEnv)

	// add driver for support yaml content
	config.AddDriver(json.Driver)
	err := config.LoadFiles("config.json")
	if err != nil {
		panic(err)
	}

	cfg := Config{}
	err = config.Decode(&cfg)

	return &cfg, nil
}
