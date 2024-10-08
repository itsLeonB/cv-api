package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewPostgresDB() *pgx.Conn {
	ctx := context.Background()

	db, err := pgx.Connect(ctx, os.Getenv("PG_URL"))
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.Ping(ctx)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return db
}
