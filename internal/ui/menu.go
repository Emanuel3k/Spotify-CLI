package ui

import "fmt"

func ShowMenu(username string) {
	fmt.Printf("Welcome to Spotify CLI, %s!\n", username)

	fmt.Println("Please select an option:")
}
