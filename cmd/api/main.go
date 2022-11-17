package main

import (
	"log"

	"github.com/SimilarEgs/tech-task/config"
	"github.com/SimilarEgs/tech-task/internal/server"
)

func main() {

	config, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	server, err := server.NewServer(config.Host + ":" + config.Port)
	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
