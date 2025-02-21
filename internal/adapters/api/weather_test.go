package api

import (
	"testing"

	"github.com/fabiowgermano/weather-api/configs"
	"github.com/stretchr/testify/assert"
)

func TestWeatherClient_GetWeather(t *testing.T) {
	config, _ := configs.LoadConfig("../../../cmd/weathersystem")
	client := NewWeatherClient(config.WeatherClientUrl, config.WeatherClientKey)

	_, err := client.GetWeather("são paulo")
	assert.NoError(t, err)
}
