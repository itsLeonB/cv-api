package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/itsLeonB/cv-api/database"
)

type Server struct {
	db     *sql.DB
	router http.Handler
}

func Init() *Server {
	s := new(Server)
	s.ConnectDatabase()
	handlers := SetupHandlers(s.db)
	s.router = SetupRouter(handlers)

	return s
}

func (s *Server) Serve() {
	defer s.db.Close()
	srv := http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: s.router,
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

func (s *Server) ConnectDatabase() {
	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatalf("fail connect db: %s", err)
	}

	s.db = db
}
