package main

import (
	"log"

	"github.com/SimilarEgs/tech-task/config"
	"github.com/SimilarEgs/tech-task/internal/server"
)

/*
Что бы изменил:
- то как мы сохраняем/удаляем/читаем юзеров (всякий раз мы считываем весь файл, что при больших объемах сущностей будем тратить много ресурсов)
- файловый репозиторий не лучший вариант
*/

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
