package weather

import "testing"

func TestWeatherCache(t *testing.T) {
	cache := Weather{
		City:        "Tokyo",
		Temperature: 25,
	}

	weatherCache.Set(
		"Tokyo",
		cache,
		0,
	)

	value, found := weatherCache.Get("Tokyo")

	if !found {
		t.Fatal("expected value in cache")
	}

	result := value.(Weather)

	if result.City != "Tokyo" {
		t.Errorf("expected Tokyo, got %s", result.City)
	}

	if result.Temperature != 25 {
		t.Errorf("expected 25, got %f", result.Temperature)
	}
}
