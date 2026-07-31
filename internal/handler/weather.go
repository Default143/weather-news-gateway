package handler

import (
	"encoding/json"
	"net/http"

	"weather-news-gateway/internal/weather"
	"weather-news-gateway/internal/geocoding"
)

// GetWeather godoc
// @Summary Get weather by city
// @Description Returns current weather information for selected city
// @Tags weather
// @Accept json
// @Produce json
// @Param city query string true "City name"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/v1/weather [get]
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
