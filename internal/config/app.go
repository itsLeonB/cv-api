package config

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsLeonB/cv-api/internal/delivery/apphttp"
	"github.com/itsLeonB/cv-api/internal/delivery/apphttp/middleware"
	"github.com/itsLeonB/cv-api/internal/delivery/apphttp/route"
	"github.com/jackc/pgx/v5"
)

type App struct {
	PostgresConn *pgx.Conn
	Router       *gin.Engine
}

func SetupApp() *App {
	conn := NewPostgresDB()
	controllers := apphttp.SetupControllers(conn)

	gin.SetMode(os.Getenv("APP_ENV"))
	r := gin.New()
	r.Use(
		gin.Recovery(),
		middleware.HandleError(),
	)

	rc := route.RouteConfig{
		Router:      r,
		Controllers: controllers,
	}

	rc.SetupRoutes()

	return &App{
		PostgresConn: conn,
		Router:       r,
	}
}

func (a *App) Serve() {
	ctx := context.Background()
	defer a.PostgresConn.Close(ctx)
	srv := http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: a.Router,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("error server listen and serve: %s", err.Error())
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("error shutting down: %e", err)
	}

	log.Println("server successfully shutdown")
}
