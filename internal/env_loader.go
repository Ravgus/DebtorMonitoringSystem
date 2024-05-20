package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	envFile := ".env.local"

	if !isFileExist(envFile) {
		envFile = ".env"
	}

	if !isFileExist(envFile) {
		fmt.Println(".env file not found!")

		os.Exit(5)
	}

	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println("Cannot load data from .env:", err)

		os.Exit(6)
	}
}

func isFileExist(envFile string) bool {
	_, err := os.Stat(envFile)

	return !os.IsNotExist(err)
}
