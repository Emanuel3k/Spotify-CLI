package main

import (
	"github.com/emanuel3k/Spotify-CLI/internal/services/auth"
	"github.com/emanuel3k/Spotify-CLI/internal/services/user"
	"github.com/emanuel3k/Spotify-CLI/internal/ui"
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

	userData, err := user.GetUserData()
	if err != nil {
		log.Fatalf("failed to get user data: %v", err)
	}
	ui.ShowMenu(userData.DisplayName)

}
