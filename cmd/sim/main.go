package main

import (
	"Regret/simulator/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found. Falling back to system environment variables.")
	}
	port := os.Getenv("SIMULATOR_PORT")
	if port == "" {
		port = ":8080"
	}
	server.SimulatorServer(port)
}
