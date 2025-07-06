package main

import (
	"github.com/emanuel3k/Spotify-CLI/internal/services/auth"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := auth.GetCredentials(); err != nil {
		log.Fatal(err)
	}
}
