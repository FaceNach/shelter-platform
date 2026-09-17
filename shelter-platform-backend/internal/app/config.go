package app

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	port        string
	databaseUrl string
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	port := os.Getenv("PORT")
	databaseUrl := os.Getenv("DATABASE_URL")

	if port == "" {
		return nil, fmt.Errorf("PORT is empty on .env")
	}

	if databaseUrl == "" {
		return nil, fmt.Errorf("DATABASE_URL is empty on .env")
	}

	return &Config{
		port:        port,
		databaseUrl: databaseUrl,
	}, nil
}
