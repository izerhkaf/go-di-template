package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ORIGINS  string
	MySQLUri string
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		ORIGINS:  os.Getenv("ALLOWED_ORIGINS"),
		MySQLUri: "", // populate from env or defaults
	}

	return cfg, nil
}

func NewConfigTest() (*Config, error) {
	_ = godotenv.Load(".env.test")

	cfg := &Config{
		ORIGINS:  os.Getenv("ALLOWED_ORIGINS"),
		MySQLUri: "", // populate from env or defaults
	}

	return cfg, nil
}
