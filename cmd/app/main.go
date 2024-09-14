package main

import (
	"log"
	server "music-exercise-tracking/internal/routes"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if _, err := os.Stat("/.dockerenv"); os.IsNotExist(err) {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
}

func main() {

	server.Start()

}
