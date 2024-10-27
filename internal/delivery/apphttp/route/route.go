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

	authRoutes := rc.Router.Group("/auth")
	authRoutes.POST("/register", rc.Controllers.Auth.HandleRegister())

	skillRoutes := rc.Router.Group("/skills")
	skillCategoryRoutes := skillRoutes.Group("/categories")
	skillCategoryRoutes.POST("", rc.Controllers.Skill.HandleInsertCategory())

	rc.Router.GET("/about", rc.Controllers.Controller.GetShortSummary())
	rc.Router.GET("/summary", rc.Controllers.Controller.HandleSummary())
}
