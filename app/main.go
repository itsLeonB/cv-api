package main

import (
	"log"

	"github.com/itsLeonB/cv-api/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading env: %e", err)
	}

	s := server.Init()
	s.Serve()
}
