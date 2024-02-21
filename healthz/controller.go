package healthz

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthzController struct {
	healthzService *HealthzService
}

func NewHealthzController(healthzService *HealthzService) *HealthzController {
	return &HealthzController{healthzService}
}

// healthCheck godoc
// @Summary      Health Checking
// @Description  Health Checking for API services
// @Produce      json
// @Success      200  {object}  HealthzResult
// @Failure      503  {object}  HealthzResult
// @Router       /healthz [get]
func (hc *HealthzController) healthCheck(ctx *gin.Context) {
	result, err := hc.healthzService.HealthCheck(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if result.Status == "ok" {
		ctx.JSON(http.StatusOK, result)
	} else {
		ctx.JSON(http.StatusServiceUnavailable, result)
	}
}

func (hc *HealthzController) RoutePattern() string {
	return "/healthz"
}

func (hc *HealthzController) RegisterControllerRoutes(rg *gin.RouterGroup) {
	rg.GET("", hc.healthCheck)
}
