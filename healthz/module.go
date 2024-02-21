package healthz

import (
	"github.com/STommydx/FolioForge/http"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(http.AsControllerRoute(NewHealthzController)),
	fx.Provide(NewHealthzService),
)
