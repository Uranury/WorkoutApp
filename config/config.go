package config

import (
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

	dsn := getEnv("DATABASE_URL", "")

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
