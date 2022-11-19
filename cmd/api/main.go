package main

import (
	"log"

	"github.com/SimilarEgs/tech-task/config"
	"github.com/SimilarEgs/tech-task/internal/server"
)

func main() {
	config_, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	api := server.NewAPI()

	server_, err := server.NewServer(config_.Host+":"+config_.Port, api)
	if err != nil {
		log.Fatal(err)
	}

	server_.Start()
}