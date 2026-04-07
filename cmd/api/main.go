package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	configuration := struct {
		DSN  string
		Port string
	}{
		DSN:  os.Getenv("DATABASE_URL"),
		Port: os.Getenv("PORT"),
	}
	_ = configuration
}
