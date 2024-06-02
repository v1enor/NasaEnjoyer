# NASA ENJOYER

## Инструкции по установке

Для установки и запуска проекта выполните следующие команды:

```bash
make build
make up
```

## Эндпоинты API

### Получение списка всех данных

Вы можете получить список всех данных, отправив GET-запрос на эндпоинт /apod.

Пример запроса:
http://localhost:9090/apod

### Получение данных по дате

Вы можете получить данные по конкретной дате, отправив GET-запрос на эндпоинт /apod/:date, где :date - это дата в формате YYYY-MM-DD.

Пример запроса:
http://localhost:9090/apod/2024-06-01