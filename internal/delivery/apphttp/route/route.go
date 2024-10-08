package route

import (
	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/cv-api/internal/delivery/apphttp"
)

type RouteConfig struct {
	Router      *gin.Engine
	Controllers *apphttp.Controllers
}

func (rc *RouteConfig) SetupRoutes() {
	rc.Router.HandleMethodNotAllowed = true
	rc.Router.ContextWithFallback = true

	rc.Router.GET("/about", rc.Controllers.Controller.GetShortSummary())
}
