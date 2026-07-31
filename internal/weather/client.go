package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

var weatherCache = cache.New(
	10*time.Minute,
	15*time.Minute,
)

type Weather struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
}

type openMeteoResponse struct {
	Current struct {
		Temperature float64 `json:"temperature_2m"`
	} `json:"current"`
}

func GetWeather(city string, lat float64, lon float64) (Weather, error) {

	if value, found := weatherCache.Get(city); found {
		log.Println("weather from cache:", city)
		return value.(Weather), nil
	}

	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current=temperature_2m",
		lat,
		lon,
	)

	log.Println("weather from api:", city)

	resp, err := httpClient.Get(url)

	if err != nil {
		return Weather{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Weather{}, fmt.Errorf("weather api returned status: %d", resp.StatusCode)
	}

	var data openMeteoResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return Weather{}, err
	}

	result := Weather{
		City:        city,
		Temperature: data.Current.Temperature,
	}

	weatherCache.Set(
		city,
		result,
		cache.DefaultExpiration,
	)

	return result, nil

}
