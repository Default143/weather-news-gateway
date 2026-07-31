package handler

import (
	"encoding/json"
	"net/http"

	"weather-news-gateway/internal/news"
)

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
