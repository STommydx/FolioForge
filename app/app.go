package app

import (
	"github.com/STommydx/FolioForge/config"
	"github.com/STommydx/FolioForge/db"
	"github.com/STommydx/FolioForge/healthz"
	"github.com/STommydx/FolioForge/http"
	"github.com/STommydx/FolioForge/logger"
	"go.uber.org/fx"
)

func New() *fx.App {
	app := fx.New(
		fx.Provide(config.NewAppConfigFromEnv),
		logger.Module,
		http.Module,
		healthz.Module,
		db.Module,
	)
	return app
}
