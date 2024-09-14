package main

import (
	"log"
	server "music-exercise-tracking/internal/routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if _, err := os.Stat("/.dockerenv"); os.IsNotExist(err) {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
	server.Start()

}
