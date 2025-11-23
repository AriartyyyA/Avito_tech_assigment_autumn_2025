# PR Reviewer Assignment Service

Сервис для автоматического управления назначением ревьюверов на Pull Request в командах разработки.

## Описание

Сервис предоставляет REST API для управления командами, пользователями и Pull Request с автоматическим назначением ревьюверов. При создании PR автоматически назначаются до двух активных ревьюверов из команды автора. Сервис поддерживает переназначение ревьюверов и массовую деактивацию пользователей команды с автоматическим переназначением открытых PR.

## Основные возможности

- **Управление командами**: создание и получение информации о командах
- **Управление пользователями**: активация/деактивация пользователей, получение статистики назначений
- **Pull Request**: создание, слияние, переназначение ревьюверов
- **Автоматическое назначение**: до 2 активных ревьюверов при создании PR
- **Массовая деактивация**: безопасное переназначение ревьюверов при деактивации команды
- **Статистика**: получение статистики по назначениям ревьюверов

## Технологии

- **Go 1.25+** - основной язык программирования
- **Gin** - веб-фреймворк
- **PostgreSQL 16** - база данных
- **Docker & Docker Compose** - контейнеризация
- **pgx/v5** - драйвер PostgreSQL

## Структура проекта

```
.
├── cmd/
│   └── main.go                 # Точка входа приложения
├── internal/
│   ├── config/                 # Конфигурация сервера
│   ├── models/                 # Модели данных
│   ├── repository/             # Слой работы с БД
│   ├── service/                # Бизнес-логика
│   └── transport/              # HTTP handlers и middleware
│       ├── dto/                # Data Transfer Objects
│       └── middleware/         # Валидация запросов
├── migrations/                 # SQL миграции (goose формат)
├── scripts/                    # Вспомогательные скрипты
├── tests/
│   └── integration/            # Интеграционные тесты
├── docker-compose.yml          # Docker Compose конфигурация
├── Dockerfile                  # Docker образ приложения
├── Makefile                    # Автоматизация задач
└── openapi.yml                 # OpenAPI спецификация
```

## Требования

- **Go** 1.25 или выше
- **Docker** и **Docker Compose**
- **PostgreSQL** 16 (или использование через Docker)

## Установка и запуск

### Быстрый старт с Docker Compose

1. Клонируйте репозиторий:
```bash
git clone <repository-url>
cd Avito_tech_assigment_autumn_2025
```

2. Запустите сервис:
```bash
make up-build
```

Или вручную:
```bash
docker-compose up -d --build
```

Сервис будет доступен на `http://localhost:8080`

Миграции применяются автоматически при запуске.

### Локальная разработка

1. Установите зависимости:
```bash
go mod download
```

2. Создайте файл `.env`:
```env
DATABASE_DSN=postgres://postgres:postgres@localhost:5432/avito_db?sslmode=disable
```

3. Запустите PostgreSQL (через Docker):
```bash
docker run -d \
  --name postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=avito_db \
  -p 5432:5432 \
  postgres:16-alpine
```

4. Примените миграции:
```bash
make db-migrate
```

5. Запустите приложение:
```bash
make run
```

Или:
```bash
go run ./cmd/main.go
```

## API Документация

Полная спецификация API доступна в файле `openapi.yml`. Для просмотра интерактивной документации используйте Swagger UI или другие инструменты.

### Основные эндпоинты

#### Команды

- `POST /team/add` - Создать команду с участниками
- `GET /team/get?team_name={name}` - Получить информацию о команде
- `GET /team/pullRequests?team_name={name}` - Получить все PR команды
- `POST /team/deactivateUsers` - Деактивировать пользователей команды

#### Пользователи

- `POST /users/setIsActive` - Изменить статус активности пользователя
- `GET /users/getReview?user_id={id}` - Получить список PR для ревью
- `GET /users/userAssignments` - Получить статистику назначений

#### Pull Request

- `POST /pullRequest/create` - Создать PR
- `POST /pullRequest/merge` - Слить PR
- `POST /pullRequest/reassign` - Переназначить ревьювера

## Тестирование

### Запуск всех тестов
```bash
make test
```

### Интеграционные тесты
```bash
make test-integration
```

### Покрытие кода
```bash
go test -cover ./...
```

## ⚡ Производительность

### Нагрузочное тестирование

Нагрузочное тестирование выполнено с использованием **k6**. Результаты показывают высокую производительность и стабильность сервиса.

#### Конфигурация теста

- **Инструмент**: k6
- **Виртуальные пользователи (VUs)**: 20
- **Длительность**: 30 секунд
- **Сценарий**: циклические запросы (создание команды, создание PR, слияние PR, получение ревью)



#### Результаты


            /\      Grafana   /‾‾/
       /\  /  \     |\  __   /  /
      /  \/    \    | |/ /  /   ‾‾\
     /          \   |   (  |  (‾)  |
    / __________ \  |_|\_\  \_____/

     execution: local
        script: k6.js
        output: -

     scenarios: (100.00%) 1 scenario, 20 max VUs, 1m0s max duration (incl. graceful stop):
              * default: 20 looping VUs for 30s (gracefulStop: 30s)

    █ TOTAL RESULTS

    checks_total.......: 12091   380.250925/s
    checks_succeeded...: 100.00% 12091 out of 12091
    checks_failed......: 0.00%   0 out of 12091

    ✓ team created or already exists
    ✓ create PR: status 201
    ✓ merge PR: status 200
    ✓ getReview: status 200

    HTTP
    http_req_duration..............: avg=11.55ms  min=93.04µs  med=4.16ms   max=2.97s    p(90)=24.6ms   p(95)=31.51ms
      { expected_response:true }...: avg=11.55ms  min=93.04µs  med=4.16ms   max=2.97s    p(90)=24.6ms   p(95)=31.51ms
    http_req_failed................: 0.00%  0 out of 12091
    http_reqs......................: 12091  380.250925/s

    EXECUTION
    iteration_duration.............: avg=124.27ms min=103.61ms med=120.61ms max=365.52ms p(90)=142.19ms p(95)=147.55ms
    iterations.....................: 4838   152.150689/s
    vus............................: 20     min=20         max=20
    vus_max........................: 20     min=20         max=20

    iterations.....................: 4838   152.150689/s
    vus............................: 20     min=20         max=20
    vus_max........................: 20     min=20         max=20
    iterations.....................: 4838   152.150689/s
    vus............................: 20     min=20         max=20
    iterations.....................: 4838   152.150689/s
    vus............................: 20     min=20         max=20
    vus_max........................: 20     min=20         max=20

    iterations.....................: 4838   152.150689/s
    vus............................: 20     min=20         max=20
    vus_max........................: 20     min=20         max=20

    NETWORK
    iterations.....................: 4838   152.150689/s
    vus............................: 20     min=20         max=20
    vus_max........................: 20     min=20         max=20

    NETWORK
    iterations.....................: 4838   152.150689/s
    vus............................: 20     min=20         max=20
    vus_max........................: 20     min=20         max=20
    iterations.....................: 4838   152.150689/s
    iterations.....................: 4838   152.150689/s
    iterations.....................: 4838   152.150689/s
    vus............................: 20     min=20         max=20
    vus_max........................: 20     min=20         max=20

    vus_max........................: 20     min=20         max=20

    vus_max........................: 20     min=20         max=20
    vus_max........................: 20     min=20         max=20

    vus_max........................: 20     min=20         max=20

    NETWORK
    vus_max........................: 20     min=20         max=20

    vus_max........................: 20     min=20         max=20

    vus_max........................: 20     min=20         max=20
    vus_max........................: 20     min=20         max=20
    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s
    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    vus_max........................: 20     min=20         max=20

    NETWORK
    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s
    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s

    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s


    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s



    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s




    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s


    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s
    vus_max........................: 20     min=20         max=20

    vus_max........................: 20     min=20         max=20
    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s




    running (0m31.8s), 00/20 VUs, 4838 complete and 0 interrupted iterations
    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s




    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s



    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s


    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s

    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s

    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    vus_max........................: 20     min=20         max=20

    vus_max........................: 20     min=20         max=20

    vus_max........................: 20     min=20         max=20
    vus_max........................: 20     min=20         max=20

    vus_max........................: 20     min=20         max=20
    vus_max........................: 20     min=20         max=20
    vus_max........................: 20     min=20         max=20

    NETWORK
    data_received..................: 1.5 GB 47 MB/s
    data_sent......................: 2.2 MB 68 kB/s



    running (0m31.8s), 00/20 VUs, 4838 complete and 0 interrupted iterations
    default ✓ [======================================] 20 VUs  30s


### Запуск нагрузочных тестов

Для запуска нагрузочных тестов используйте k6:

```bash
k6 run k6.js
```
## Makefile команды

```bash
# Docker операции
make build          # Собрать Docker образы
make up             # Запустить сервисы
make up-build       # Собрать и запустить
make down           # Остановить сервисы
make down-clean     # Остановить и удалить volumes
make restart        # Перезапустить сервисы
make logs           # Просмотр логов

# Разработка
make run            # Запустить приложение локально
make dev            # Режим разработки

# Тестирование
make test           # Запустить все тесты
make test-integration # Интеграционные тесты

# Code Quality
make lint           # Запустить линтер
make fmt            # Форматировать код

# База данных
make db-connect     # Подключиться к БД
make db-migrate     # Применить миграции
```

## Миграции

Миграции находятся в директории `migrations/` и используют формат goose. При запуске через Docker Compose они применяются автоматически.

Для ручного применения:
```bash
make db-migrate
```

## Примеры использования

### Создание команды

```bash
curl -X POST http://localhost:8080/team/add \
  -H "Content-Type: application/json" \
  -d '{
    "team_name": "backend",
    "members": [
      {"user_id": "u1", "username": "Alice", "is_active": true},
      {"user_id": "u2", "username": "Bob", "is_active": true}
    ]
  }'
```

### Создание Pull Request

```bash
curl -X POST http://localhost:8080/pullRequest/create \
  -H "Content-Type: application/json" \
  -d '{
    "pull_request_id": "pr-1001",
    "pull_request_name": "Add search feature",
    "author_id": "u1"
  }'
```

### Переназначение ревьювера

```bash
curl -X POST http://localhost:8080/pullRequest/reassign \
  -H "Content-Type: application/json" \
  -d '{
    "pull_request_id": "pr-1001",
    "old_user_id": "u2"
  }'
```

### Деактивация команды

```bash
curl -X POST http://localhost:8080/team/deactivateUsers \
  -H "Content-Type: application/json" \
  -d '{
    "team_name": "backend"
  }'
```

## Переменные окружения

- `DATABASE_DSN` - строка подключения к PostgreSQL
  - Формат: `postgres://user:password@host:port/database?sslmode=disable`
  - По умолчанию (Docker): `postgres://postgres:postgres@postgres:5432/avito_db?sslmode=disable`

## Docker

### Сервисы

- **pr_service_db** - PostgreSQL база данных
- **pr_service_migrate** - Применение миграций
- **pr_service_app** - Основное приложение

### Порты

- `8080` - HTTP API
- `5432` - PostgreSQL

