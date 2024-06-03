# NASA ENJOYER

## Окружение
Сначало необходимо настроить окружение.
Установить свои данные, или оставить по умолчанию в example.env и переименовать в .env 

## Установка
Для установки и запуска проекта выполните следующие команды:

```bash
make build
make up
```

## Эндпоинты API

### Получение списка всех данных

Вы можете получить список всех данных, отправив GET-запрос на эндпоинт /apod.

Пример запроса (по умолчанию):
http://localhost:9090/apod

### Получение данных по дате

Вы можете получить данные по конкретной дате, отправив GET-запрос на эндпоинт /apod/:date, где :date - это дата в формате YYYY-MM-DD.

Пример запроса  (по умолчанию):
http://localhost:9090/apod/2024-06-01
