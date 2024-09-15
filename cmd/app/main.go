package main

import (
	"log"
	server "music-exercise-tracking/internal/routes"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(dir, ".env")
	err = godotenv.Load(environmentPath)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	server.Start()

}
