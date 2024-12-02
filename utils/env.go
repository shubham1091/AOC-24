package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetSessionCookie() (string, error) {
	if os.Getenv("SESSION_COOKIE") == "" {
		return "", fmt.Errorf("SESSION_COOKIE environment variable not set")
	}
	return os.Getenv("SESSION_COOKIE"), nil
}
