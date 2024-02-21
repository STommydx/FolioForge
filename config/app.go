package config

import (
	"context"

	"github.com/STommydx/FolioForge/db"
	"github.com/STommydx/FolioForge/http"
	"github.com/STommydx/FolioForge/logger"
	"github.com/sethvargo/go-envconfig"
	"go.uber.org/fx"
)

type AppConfig struct {
	fx.Out
	Postgres *db.Config     `env:", prefix=POSTGRES_"`
	Logger   *logger.Config `env:", prefix=LOG_"`
	Api      *http.Config   `env:", prefix=API_"`
}

func NewAppConfigFromEnv() (AppConfig, error) {
	ctx := context.Background()
	config := AppConfig{}
	if err := envconfig.Process(ctx, &config); err != nil {
		return config, err
	}
	return config, nil
}
