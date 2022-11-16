package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct - contains config of the program
type Config struct {
	HTTP
}

// HTTP struct - contains config of the HTTP server
type HTTP struct {
	Host string
	Port string
}

// ParseConfig func - returns parsed config struct which was read from the provided config file
func ParseConfig() (*Config, error) {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("[Error] .env file didn't load: %s", err.Error())
	}

	var c Config

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort != "" {
		c.HTTP.Port = httpPort
	}

	httpHost := os.Getenv("HTTP_HOST")
	if httpHost != "" {
		c.Host = httpHost
	}

	return &c, nil
}
