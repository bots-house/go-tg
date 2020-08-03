# Birzzha API

## Flow

[Наш подход к работе с GitHub](https://www.notion.so/linch/GitHub-47e8ce54dda4417d89955617365e5859)

## Настройка окружения

В проект используется приватные зависимости, для того чтобы они подтянулись с GitHub по ssh, нужно выполнить команду:

```bash
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

Так же нужно установить Go и docker-compose.

После того как проект будет склонен локально, нужно скопировать `.env.example`, как `.env` и прописать упущенные значения: 
  - `BRZ_BOT_TOKEN`: токен бота, нужно создать его в [BotFather](https://t.me/BotFather)
  - `BRZ_BOT_WEBHOOK_DOMAIN`: домен на который будет установлен webhook от Telegram и Interkassa. Можно получить запустив `ngrok http 8000`

## Запуск

- запуск сервера и воркера в одном процессе:  `make run`
- запуск сервера: `make run-server`
- запуск воркера: `make run-worker`

## Инструменты

1. Установите jq. 
  - MacOS: `brew install jq`;
  - Ubuntu/Debian: `sudo apt-get install jq`; 

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
