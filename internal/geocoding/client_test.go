package geocoding

import "testing"

func TestLocationCache(t *testing.T) {
	location := Location{
		Latitude:  35.6762,
		Longitude: 139.6503,
	}

	locationCache.Set(
		"Tokyo",
		location,
		0,
	)

	value, found := locationCache.Get("Tokyo")

	if !found {
		t.Fatal("expected location in cache")
	}

	result := value.(Location)

	if result.Latitude != 35.6762 {
		t.Errorf("expected latitude 35.6762, got %f", result.Latitude)
	}

	if result.Longitude != 139.6503 {
		t.Errorf("expected longitude 139.6503, got %f", result.Longitude)
	}
}
