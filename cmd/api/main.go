package main

import (
	"log"

	"github.com/SimilarEgs/tech-task/config"
	"github.com/SimilarEgs/tech-task/internal/server"
	"github.com/SimilarEgs/tech-task/internal/service"
)

func main() {

	config, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	userService := service.NewUserService()

	api := server.NewAPI(*service.NewServerService(userService))

	server, err := server.NewServer(config.Host+":"+config.Port, api)
	if err != nil {
		log.Fatal(err)
	}

	server.Start()
}
