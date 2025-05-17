package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Env          string
	Port         string
	GeminiApiKey string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Env:          getEnv("ENV", "dev"),
		Port:         getEnv("PORT", ":8080"),
		GeminiApiKey: getEnv("GEMINI_API_KEY", "apiKey"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvOrError(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	panic(fmt.Sprintf("Environment variable %s is not set", key))

}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}

		return b
	}

	return fallback
}
