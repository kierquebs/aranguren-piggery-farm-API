package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Env func to get env value
func GetEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
