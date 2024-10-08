package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectPostgres() (*sql.DB, error) {
	url := os.Getenv("PG_URL")

	db, err := sql.Open("pgx", url)
	if err != nil {
		log.Printf("error on ConnectPostgres(): %s\ntype: %T\ndetails: %v\n", err.Error(), err, err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("error on ConnectPostgres(): %s\ntype: %T\ndetails: %v\n", err.Error(), err, err)
		return nil, err
	}

	return db, nil
}
