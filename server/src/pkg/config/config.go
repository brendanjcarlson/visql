package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func MustLoadEnv(filenames ...string) {
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Fatalf("error loading env files: %v\n", err.Error())
	}
}

func MustGet(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("error getting env variable: %v\n", key)
	}

	return value
}

func GetOrDefault(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return value
}
