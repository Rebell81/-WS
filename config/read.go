package config

import (
	"context"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
)

func ReadConfig(ctx context.Context) (*Config, error) {
	loader := confita.NewLoader(env.NewBackend(), file.NewOptionalBackend("config.yaml"))

	cfg := Config{}

	err := loader.Load(ctx, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
