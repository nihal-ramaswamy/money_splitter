package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetDotEnvVariable(key string) string {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
