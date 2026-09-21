package main

import (
	"Regret/src/api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found. Falling back to system environment variables.")
	}

	port := os.Getenv("API_PORT")

	if port == "" {
		port = ":8081"
	}
	api.MainServer(port)
}
