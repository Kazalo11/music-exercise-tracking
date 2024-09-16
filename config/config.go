package config

import (
	"os"
)

func GetFrontendHost() string {
	if os.Getenv("ENV") == "prod" {
		return "https://frontend-1052978901140.europe-west2.run.app"
	}
	return "http://localhost"
}
