package main

import (
	"log"
	"os"
)

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		log.Fatal("Missing required environment variable HOST")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Missing required environment variable PORT")
	}

	StartService(host, port)
}
