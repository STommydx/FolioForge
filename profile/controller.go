package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileService *ProfileService
}

func NewProfileController(profileService *ProfileService) *ProfileController {
	return &ProfileController{profileService}
}

func (pc *ProfileController) getProfileByName(ctx *gin.Context) {
	profileName := ctx.Param("profile_name")
	profile, err := pc.profileService.GetProfileByName(ctx, profileName)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if profile == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, profile)
}

func (pc *ProfileController) RoutePattern() string {
	return "/profiles"
}

func (pc *ProfileController) RegisterControllerRoutes(rg *gin.RouterGroup) {
	rg.GET("/:profile_name", pc.getProfileByName)
}
