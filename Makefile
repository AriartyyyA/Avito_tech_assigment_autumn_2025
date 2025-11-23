.PHONY: help build up down restart logs clean test lint fmt deps db-connect db-migrate dev run

DOCKER_COMPOSE = docker-compose
APP_NAME = pr_service_app
POSTGRES_NAME = pr_service_db
MIGRATE_NAME = pr_service_migrate

build:
	$(DOCKER_COMPOSE) build

up:
	$(DOCKER_COMPOSE) up -d

up-build:
	$(DOCKER_COMPOSE) up -d --build

down:
	$(DOCKER_COMPOSE) down

down-clean:
	$(DOCKER_COMPOSE) down -v

restart: 
	$(DOCKER_COMPOSE) restart

run:
	@if [ ! -f .env ]; then \
		echo "Creating .env file from example..."; \
		echo "DATABASE_DSN=postgres://postgres:postgres@localhost:5432/avito_db?sslmode=disable" > .env; \
	fi
	go run ./cmd/main.go

test:
	go test -v ./...

test-integration:
	go test -v ./tests/integration/...

lint:
	@if ! command -v golangci-lint > /dev/null; then \
		echo "Installing golangci-lint..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	golangci-lint run


