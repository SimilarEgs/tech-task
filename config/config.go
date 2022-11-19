package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP
}

type HTTP struct {
	Host string
	Port string
}

func ParseConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("[Error] .env file didn't load: %s", err.Error())
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
