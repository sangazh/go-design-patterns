package facade

import (
	"testing"
)

func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := new(CurrentWeatherData)
	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}

	if weather.ID != 2643743 {
		t.Error("London id is 2643743, got: %d", weather.ID)
	}
}
