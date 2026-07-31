# Weather News Gateway

![Go tests](https://github.com/Default143/weather-news-gateway/actions/workflows/tests.yml/badge.svg)

REST API на Go, который объединяет данные о погоде и новости из внешних сервисов.

## Возможности

- Получение погоды по городу
- Поиск координат через геокодирование
- Получение новостей по теме
- Кэширование запросов
- Swagger-документация
- Docker-запуск

## Технологии

- Go
- Chi Router
- Open-Meteo API
- NewsAPI
- Docker
- Swagger/OpenAPI

## Запуск

Создать файл `.env`:

NEWS_API_KEY=your_key

Запуск:

go run ./cmd/gateway

Swagger:

http://localhost:8080/swagger/index.html

## API

GET /api/v1/weather?city=Tokyo

GET /api/v1/news?topic=apple

GET /api/v1/dashboard?city=Tokyo&topic=apple

## Скриншоты

## Swagger UI

![Swagger](screenshots/Swagger UI.png)

## Dashboard API response

![Dashboard](screenshots/dasboard-response.png)
