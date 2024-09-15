package main

import (
	server "music-exercise-tracking/internal/routes"

	"github.com/joho/godotenv"
)

func init() {
	path := ".env"
	for {
		err := godotenv.Load(path)
		if err == nil {
			break
		}
		path = "../" + path
	}
}

func main() {

	server.Start()

}
