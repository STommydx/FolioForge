package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Config struct {
	ListenAddr string `env:"LISTEN_ADDR, default=:8080"`
}

func newHttp(config *Config, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    config.ListenAddr,
		Handler: router,
	}
}

func runHttpServer(lifecycle fx.Lifecycle, httpServer *http.Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go httpServer.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return httpServer.Shutdown(ctx)
		},
	})
}
