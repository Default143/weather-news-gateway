## Weather News Gateway

API-шлюз на Go, который агрегирует данные о погоде и новостях.

## Возможности

* получение погоды по названию города;
* получение новостей по теме;
* объединённый endpoint `/dashboard`;
* кэширование запросов;
* Swagger/OpenAPI документация;
* Docker и Docker Compose.

## Запуск

docker compose up --build

## Swagger

Открыть:

http://IP_WSL:8080/swagger/index.html

## Примеры запросов

### Погода

curl "http://localhost:8080/api/v1/weather?city=Tokyo"

### Новости

curl "http://localhost:8080/api/v1/news?topic=apple"

### Dashboard

curl "http://localhost:8080/api/v1/dashboard?city=Tokyo&topic=apple"


