// pkg/config/config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIToken           string
	DBConnectionString string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		APIToken:           os.Getenv("API_TOKEN"),
		DBConnectionString: os.Getenv("DB_CONNECTION_STRING"),
	}
}
