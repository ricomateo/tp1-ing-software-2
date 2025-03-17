package main

import (
	"log"
)

func main() {
	logger := log.Default()

	host, err := GetEnvVar("HOST")
	if err != nil {
		logger.Print("Error: ", err)
	}
	port, err := GetEnvVar("PORT")
	if err != nil {
		logger.Print("Error: ", err)
	}
	environment, err := GetEnvVar("ENVIRONMENT")
	if err != nil {
		logger.Print("Error: ", err)
	}

	logger.Println("Enviroment: ", environment)
	StartService(host, port)
}
