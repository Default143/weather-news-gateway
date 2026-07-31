package handler

import (
	"encoding/json"
	"net/http"

	"weather-news-gateway/internal/news"
)

// GetNews godoc
// @Summary Get news by topic
// @Description Returns news articles by search topic
// @Tags news
// @Accept json
// @Produce json
// @Param topic query string true "News topic"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/v1/news [get]
func News(w http.ResponseWriter, r *http.Request) {

	topic := r.URL.Query().Get("topic")

	result, err := news.GetNews(topic)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)
}
