package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func loadConfig(env string) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	configFile := fmt.Sprintf("config.%s.yaml", env)
	viper.SetConfigName(configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading in environment variables")
	}

}

func GetConfig() {
	env := os.Getenv("ENV")
	if env != "prod" {
		env = "dev"
	}
	loadConfig(env)
}
