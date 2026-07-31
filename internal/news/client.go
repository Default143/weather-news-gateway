package news

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type apiResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

var newsCache = cache.New(
	10*time.Minute,
	15*time.Minute,
)

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

func GetNews(topic string) ([]Article, error) {

	if value, found := newsCache.Get(topic); found {
		log.Println("news from cache:", topic)
		return value.([]Article), nil
	}

	apiKey := os.Getenv("NEWS_API_KEY")

	if apiKey == "" {
		return nil, fmt.Errorf("NEWS_API_KEY is not set")
	}

	url := fmt.Sprintf(
		"https://newsapi.org/v2/everything?q=%s&pageSize=5&apiKey=%s",
		topic,
		apiKey,
	)

	log.Println("news from api:", topic)

	resp, err := httpClient.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("news api returned status: %d", resp.StatusCode)
	}

	var data apiResponse

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	newsCache.Set(
		topic,
		data.Articles,
		cache.DefaultExpiration,
	)

	return data.Articles, nil
}
