# Birzzha API

## Разработка
- Запуск приложения:
  - `make`
- Слой [DAL](https://en.wikipedia.org/wiki/Data_access_layer) генерируется с помощью [sqlboiler](https://github.com/volatiletech/sqlboiler).
  - `make generate-dal`
- Слой API генерируется с помощью [go-swagger](https://github.com/go-swagger/go-swagger).
  - `make generate-api`
- Во избежения использования разработчиками разных версий кодогенераторов, при первом вызове они будут установлены в `.bin/`.
