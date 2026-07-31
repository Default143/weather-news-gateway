package handler

import (
	"encoding/json"
	"net/http"

	"weather-news-gateway/internal/weather"
	"weather-news-gateway/internal/geocoding"
)

func Weather(w http.ResponseWriter, r *http.Request) {

	city := r.URL.Query().Get("city")

	location, err := geocoding.GetCoordinates(city)

	if err != nil {
		http.Error(w, "city not found", http.StatusBadRequest)
		return
	}

	result, err := weather.GetWeather(
		city,
		location.Latitude,
		location.Longitude,
	)

	if err != nil {
		http.Error(w, "weather service unavailable", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)
}
