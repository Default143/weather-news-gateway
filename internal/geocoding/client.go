package geocoding

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"time"
)

var locationCache = cache.New(
	10*time.Minute,
	15*time.Minute,
)

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

type Location struct {
	Latitude  float64
	Longitude float64
}

type response struct {
	Results []struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"results"`
}

func GetCoordinates(city string) (Location, error) {

	if value, found := locationCache.Get(city); found {
		log.Println("location from cache:", city)
		return value.(Location), nil
	}

	url := fmt.Sprintf(
		"https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1",
		city,
	)

	resp, err := httpClient.Get(url)

	if err != nil {
		return Location{}, err
	}

	defer resp.Body.Close()

	var data response

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		return Location{}, err
	}

	if len(data.Results) == 0 {
		return Location{}, fmt.Errorf("city not found")
	}

	result := Location{
		Latitude:  data.Results[0].Latitude,
		Longitude: data.Results[0].Longitude,
	}

	locationCache.Set(
		city,
		result,
		cache.DefaultExpiration,
	)

	log.Println("location from api:", city)

	return result, nil
}
