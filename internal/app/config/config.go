package config

import (
	"github.com/caarlos0/env/v6"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string `env:"BOT_TOKEN"`
}

func New() (*Config, error) {
	err := loadEnv()
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func loadEnv() error {
	err := godotenv.Load(os.Getenv("ENV_FILE"))
	if err != nil {
		return err
	}
	return nil
}
