package handler

import (
	"encoding/json"
	"net/http"

	"weather-news-gateway/internal/geocoding"
	"weather-news-gateway/internal/news"
	"weather-news-gateway/internal/weather"
)

type DashboardResponse struct {
	Weather weather.Weather `json:"weather"`
	News    []news.Article  `json:"news"`
}

// Dashboard godoc
// @Summary Get weather and news
// @Description Returns weather and news for a city and topic
// @Tags dashboard
// @Produce json
// @Param city query string true "City name"
// @Param topic query string true "News topic"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/v1/dashboard [get]
func Dashboard(w http.ResponseWriter, r *http.Request) {

	city := r.URL.Query().Get("city")
	topic := r.URL.Query().Get("topic")

	if city == "" {
		WriteError(w, http.StatusBadRequest, "city is required")
		return
	}

	if topic == "" {
		WriteError(w, http.StatusBadRequest, "topic is required")
		return
	}

	location, err := geocoding.GetCoordinates(city)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	currentWeather, err := weather.GetWeather(
		city,
		location.Latitude,
		location.Longitude,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	articles, err := news.GetNews(topic)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := DashboardResponse{
		Weather: currentWeather,
		News:    articles,
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	encoder.Encode(response)
}
