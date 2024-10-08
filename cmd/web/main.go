package main

import (
	"log"

	"github.com/itsLeonB/cv-api/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading env: %e", err)
	}

	a := config.SetupApp()
	a.Serve()
}
