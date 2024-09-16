package config

import (
	"os"
)

func GetBackendHost() string {
	if os.Getenv("ENV") == "prod" {
		return "https://backend-1052978901140.europe-west2.run.app"
	}
	return "http://localhost"
}

func GetDomain() string {
	if os.Getenv("ENV") == "prod" {
		return ".europe-west2.run.app"
	}
	return "localhost"
}

func GetFrontendUrl() string {
	if os.Getenv("ENV") == "prod" {
		return "https://frontend-1052978901140.europe-west2.run.app"
	}
	return "http://localhost:3000"
}

func IsSecure() bool {
	return os.Getenv("ENV") == "prod"
}
