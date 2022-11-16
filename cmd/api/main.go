package main

import (
	"fmt"
	"log"

	"github.com/SimilarEgs/tech-task/config"
	"github.com/SimilarEgs/tech-task/internal/server"
)

func main() {

	fmt.Println("[Info] starting server")

	// load config
	config, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}
	server, err := server.NewServer(config.Host + ":" + config.Port)
	server.ListenAndServe()
}
