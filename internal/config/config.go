package config

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	OllamaURL string
	APIKeys   []string
}

func Load() (*Config, error) {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ollamaURL := os.Getenv("OLLAMA_URL")
	if ollamaURL == "" {
		return nil, errors.New("OLLAMA_URL is required")
	}

	apiKeysRaw := os.Getenv("API_KEYS")
	if apiKeysRaw == "" {
		return nil, errors.New("API_KEYS is required")
	}

	keys := strings.Split(apiKeysRaw, ",")
	for i, k := range keys {
		keys[i] = strings.TrimSpace(k)
	}

	return &Config{
		Port:      port,
		OllamaURL: ollamaURL,
		APIKeys:   keys,
	}, nil
}
