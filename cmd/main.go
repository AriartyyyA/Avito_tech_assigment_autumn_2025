package main

import (
	"log"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/config"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/repository"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/service"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport"
)

func main() {
	repo := repository.NewRepository()

	service := service.NewService(repo)

	handlers := transport.NewHandler(service)

	srv := new(config.Server)

	if err := srv.Run(handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
	log.Println("server is starting")
}
