package main

import (
	"gyme_backend/app"
	"gyme_backend/config"
	"log"

	"github.com/joho/godotenv"
)
func main() {
    err := godotenv.Load("config/config.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    config.ConnectDatabase()
    app.Run()
}
