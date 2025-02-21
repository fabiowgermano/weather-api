package usecase

import (
	"errors"
	"testing"

	"github.com/fabiowgermano/weather-api/internal/adapters/api"
	"github.com/fabiowgermano/weather-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type LocationClientMock struct {
	mock.Mock
}

func (m *LocationClientMock) GetLocation(cep entity.CEP) (string, error) {
	args := m.Called(cep)
	if cep == "13180000" {
		return "Sumaré", nil
	} else {
		return "", args.Error(1)
	}
}

type WeatherClientMock struct {
}

func (m *WeatherClientMock) GetWeather(city string) (float64, error) {
	return 28.5, nil
}

func TestGetTempUseCase(t *testing.T) {
	locationClient := &LocationClientMock{}
	locationClient.On("GetLocation", entity.CEP("13180000")).Return("Sumaré", nil)
	locationClient.On("GetLocation", entity.CEP("13180022")).Return("", api.ErrNotFoundZipcode)
	locationClient.On("GetLocation", entity.CEP("1318000")).Return("", errors.New("bad request"))

	weatherClient := &WeatherClientMock{}

	getTemp := NewGetTempUseCase(locationClient, weatherClient)

	expected := TempOutput{
		TempC: 28.5,
		TempF: 83.3,
		TempK: 301.5,
	}

	dto, err := getTemp.Execute("13180000")
	assert.Equal(t, expected, dto)
	assert.Nil(t, err)

	_, err = getTemp.Execute("13180022")
	assert.Equal(t, api.ErrNotFoundZipcode, err)

	_, err = getTemp.Execute("1318000")
	assert.Equal(t, entity.ErrInvalidZipcode, err)
}
