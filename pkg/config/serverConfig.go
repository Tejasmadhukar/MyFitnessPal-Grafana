package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT          = "80"
	HOST_URL      = "http://localhost"
	ASSETS_DIR    = "internal/assets/"
	GRAFANA_TOKEN = "*"
	GRAFANA_HOST  = "http://localhost:3000"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Could not read env file")
	}

	if val, ok := os.LookupEnv("PORT"); ok {
		PORT = val
	}

	if val, ok := os.LookupEnv("ASSETS_DIR"); ok {
		ASSETS_DIR = val
	}

	if val, ok := os.LookupEnv("GRAFANA_TOKEN"); ok {
		GRAFANA_TOKEN = val
	}

	if val, ok := os.LookupEnv("GRAFANA_HOST"); ok {
		GRAFANA_HOST = val
	}

	if val, ok := os.LookupEnv("HOST_URL"); ok {
		HOST_URL = val
	}
}
