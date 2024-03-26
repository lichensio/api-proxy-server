package main

import (
	"api_proxy_server/pkg/db"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load() // Load .env file
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Construct the DSN from .env variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"))

	// Initialize the repository
	repo, err := db.NewRepository(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Here you would initialize your server and pass the repository to it
	fmt.Println("Repository initialized successfully", repo)
	// Additional setup for your server...
}
