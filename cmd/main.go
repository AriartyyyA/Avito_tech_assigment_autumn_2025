package main

import (
	"context"
	"log"
	"os"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/config"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/repository"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/service"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_DSN"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err.Error())
	}
	defer conn.Close()

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("Error ping DB")
	}

	repo := repository.NewRepository(conn)

	service := service.NewService(repo)

	handlers := transport.NewHandler(service)

	srv := new(config.Server)

	if err := srv.Run(handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
	log.Println("server is starting")
}
