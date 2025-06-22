package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	JWTSecret string
	DbUrl     string
}

func Load() *Config {
	_ = loadEnv()

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		getEnv("POSTGRES_USER", ""),
		getEnv("POSTGRES_PASSWORD", ""),
		getEnv("POSTGRES_DB", "workoutapp"),
		getEnv("DATABASE_HOST", "localhost"),
		getEnv("DATABASE_PORT", "5434"),
		getEnv("SSL_MODE", "disable"),
	)

	cfg := &Config{
		Port:      getEnv("PORT", "4040"),
		DbUrl:     dsn,
		JWTSecret: getEnv("JWT_SECRET", ""),
	}

	if cfg.DbUrl == "" {
		log.Fatal("DATABASE_URL is required")
	}

	return cfg
}

func loadEnv() error {
	return godotenv.Load()
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
