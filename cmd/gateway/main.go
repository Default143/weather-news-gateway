// @title Weather News Gateway API
// @version 1.0
// @description API gateway for weather and news aggregation

package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "weather-news-gateway/docs"

	"weather-news-gateway/internal/config"
	"weather-news-gateway/internal/handler"
)

func main() {

	cfg, err := config.Load()

	if err != nil {
		panic(err)
	}

	_ = cfg

	godotenv.Load()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/health", handler.Health)
	r.Get("/api/v1/weather", handler.Weather)
	r.Get("/api/v1/news", handler.News)
	r.Get("/api/v1/dashboard", handler.Dashboard)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("server started on :8080")

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
