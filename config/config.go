package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	Host             string
	Port             string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	HTTPPort         string
	HTTPHost         string
}

func Load() (Config, error) {
	return Config{
		Host:             os.Getenv("HOST"),
		Port:             os.Getenv("PORT"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		HTTPPort:         os.Getenv("8080"),
		HTTPHost:         os.Getenv("localhost"),
	}, nil
}
