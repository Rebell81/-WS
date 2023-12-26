package config

import (
	"context"
	"fmt"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
)

func ReadConfig(ctx context.Context) (*Config, error) {
	loader := confita.NewLoader(env.NewBackend(), file.NewOptionalBackend("config.json"))

	cfg := Config{
		DurationSeconds: 60,
		ManualCheckOnly: false,
		SSL:             false,
		RutrackerHost:   "https://api.rutracker.cc",
	}

	err := loader.Load(ctx, &cfg)
	if err != nil {
		return nil, err
	}

	fmt.Println(cfg)

	return &cfg, nil
}
