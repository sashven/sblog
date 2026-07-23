package config

import "os"

type Config struct {
	Env        string
	HTTPAddr   string
	ContentDir string
}

func Load() (Config, error) {
	return Config{
		Env:        getEnv("APP_ENV", "development"),
		HTTPAddr:   getEnv("HTTP_ADDR", ":8080"),
		ContentDir: getEnv("CONTENT_DIR", "content/posts"),
	}, nil
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
