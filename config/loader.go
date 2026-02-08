package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

func Load(ctx context.Context) (AppConfig, error) {
	cfg := new(AppConfig)
	if err := envconfig.Process(ctx, cfg); err != nil {
		return AppConfig{}, err
	}

	return *cfg, nil
}
