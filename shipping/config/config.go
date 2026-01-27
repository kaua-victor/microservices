package config

import (
	"log"
	"os"
	"strconv"
)

func GetApplicationPort() int {
	port, err := strconv.Atoi(os.Getenv("APPLICATION_PORT"))
	if err != nil {
		log.Fatal("APPLICATION_PORT missing")
	}
	return port
}

func GetDataSourceURL() string {
	url := os.Getenv("DATA_SOURCE_URL")
	if url == "" {
		log.Fatal("DATA_SOURCE_URL missing")
	}
	return url
}
