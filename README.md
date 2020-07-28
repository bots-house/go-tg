# Birzzha API

## Настройка окружения

В проект используется приватные зависимости, для того чтобы они подтянулись с GitHub по ssh, нужно выполнить команду:

```bash
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

Так же нужно установить Go и docker-compose.

После того как проект будет склонен локально, нужно скопировать `.env.example`, как `.env` и прописать упущенные значения.

## Запуск

Для начала нужно поднять зависимости:

```bash
docker-compose up -d
```

После, можно запускать сервис состоящий из HTTP-сервера и cron-like воркера.

```
go run main.go -config .env.local
```

Для запуска только сервера: `go run main.go -config .env.local -server`
Для запуска только воркера: `go run main.go -config .env.local -worker`


## Инструменты

### Генерация слоя DAL

```bash
make generate-dal
```

### Генерация слоя REST API

```bash
make generate-api
```

### Линтер

```bash
make lint
```

### Подключится к Postgres

```bash
make psql
```

### Подключится к Redis

```bash
make redis-cli
```
