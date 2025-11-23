package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/config"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/repository"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/service"
	"github.com/AriartyyyA/Avito_tech_assigment_autumn_2025/internal/transport"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_DSN"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err.Error())
	}
	defer func() {
		log.Println("closing db connection pool")
		conn.Close()
	}()

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("Error ping DB")
	}

	repo := repository.NewRepository(conn)

	service := service.NewService(repo)

	handlers := transport.NewHandler(service)

	srv := new(config.Server)

	serverErrChan := make(chan error, 1)

	go func() {
		if err := srv.Run(handlers.InitRoutes()); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("http server ListenAndServe error: %v", err)
			serverErrChan <- err
			return
		}

		close(serverErrChan)
	}()

	log.Println("server is starting")

	select {
	case <-ctx.Done():
		log.Println("shutdown signal")
	case err := <-serverErrChan:
		if err != nil {
			log.Printf("server run error: %v", err)
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("server forced to shutdown: %v", err)
	} else {
		log.Println("server exited gracefully")
	}
}
