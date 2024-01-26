package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT             = "80"
	ASSETS_DIR       = "internal/assets/"
	GRAFANA_USER_ID  = "admin"
	GRAFANA_PASSWORD = "admin"
	GRAFANA_HOST     = "http://localhost:3000"
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

	if val, ok := os.LookupEnv("GRAFANA_USER_ID"); ok {
		GRAFANA_USER_ID = val
	}

	if val, ok := os.LookupEnv("GRAFANA_PASSWORD"); ok {
		GRAFANA_PASSWORD = val
	}

	if val, ok := os.LookupEnv("GRAFANA_HOST"); ok {
		GRAFANA_HOST = val
	}

}
