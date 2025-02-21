package api

import (
	"strings"
	"testing"

	"github.com/fabiowgermano/weather-api/configs"
	"github.com/stretchr/testify/assert"
)

func TestLocationClient_GetLocation(t *testing.T) {
	config, _ := configs.LoadConfig("../../../cmd/weathersystem")
	client := NewLocationClient(config.LocationClientUrl)

	city, err := client.GetLocation("13180000")
	assert.Equal(t, "SUMARÉ", strings.ToUpper(city))
	assert.NoError(t, err)

	_, err = client.GetLocation("")
	assert.Error(t, err, "invalid zipcode")

	_, err = client.GetLocation("1318002")
	assert.Error(t, err, "not exists zipcode")
}
