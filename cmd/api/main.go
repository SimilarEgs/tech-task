package main

import (
	"fmt"

	"github.com/SimilarEgs/tech-task/config"
)

func main() {

	fmt.Println("[Info] starting server")

	config, err := config.ParseConfig()

}
