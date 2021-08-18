package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ReadEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
