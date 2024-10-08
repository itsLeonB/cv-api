package server

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/cv-api/handler"
	"github.com/itsLeonB/cv-api/repository"
	"github.com/itsLeonB/cv-api/service"
)

type Handlers struct {
	// rootHandler    *handler.RootHandler
	appHandler *handler.Handler
}

func SetupHandlers(db *sql.DB) *Handlers {
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	handler := handler.NewHandler(svc)

	return &Handlers{
		appHandler: handler,
	}
}

func SetupRouter(handlers *Handlers) http.Handler {
	gin.SetMode(os.Getenv("APP_ENV"))
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.ContextWithFallback = true
	// r.Use(gin.Recovery(), middleware.Error(), middleware.Timeout(), middleware.CORS())

	// r.NoRoute(handlers.rootHandler.NotFound())
	// r.GET("/", handlers.rootHandler.Root())

	r.GET("/about", handlers.appHandler.GetShortSummary())

	return r
}
