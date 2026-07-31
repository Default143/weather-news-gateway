// @title Weather News Gateway API
// @version 1.0
// @description API gateway for weather and news aggregation

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	go func() {
		err = server.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-stop

	log.Println("stopping server")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("server stopped")
}
