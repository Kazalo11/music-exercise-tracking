package main

import (
	"log"
	server "music-exercise-tracking/internal/routes"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {

	server.Start()

}
