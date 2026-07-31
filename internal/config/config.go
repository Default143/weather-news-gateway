package config

import (
	"fmt"
	"os"
)

type Config struct {
	NewsAPIKey string
}

func Load() (Config, error) {
	key := os.Getenv("NEWS_API_KEY")

	if key == "" {
		return Config{}, fmt.Errorf("NEWS_API_KEY is not set")
	}

	return Config {
		NewsAPIKey: key,
	}, nil
}
