package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
	ENV  string
	DATABASE_URL string
}

func MustLoad() Config {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT is required")
	}

	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is required")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		panic("DATABASE_URL is required")
	}

	return Config {
		PORT: port,
		ENV: env,
		DATABASE_URL: dbUrl,
	}
}