package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GeminiAPIKey string `envconfig:"GEMINI_API_KEY" required:"true"`
	OutputDir    string `envconfig:"OUTPUT_DIR" default:"./generated_images"`
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &cfg, nil
}
