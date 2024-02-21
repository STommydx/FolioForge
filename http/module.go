package http

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(newHttp),
	fx.Invoke(runHttpServer),
	fx.Provide(NewRouter),
	fx.Provide(NewSwaggerHandler),
)
