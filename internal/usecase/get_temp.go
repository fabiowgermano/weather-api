package usecase

import (
	"math"

	"github.com/fabiowgermano/weather-api/internal/entity"
)

type TempOutput struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type GetTempUseCase struct {
	LocationClient entity.LocationClientInterface
	WeatherClient  entity.WeatherClientInterface
}

func NewGetTempUseCase(
	locationClient entity.LocationClientInterface,
	weatherClient entity.WeatherClientInterface,
) *GetTempUseCase {
	return &GetTempUseCase{
		LocationClient: locationClient,
		WeatherClient:  weatherClient,
	}
}

func (g *GetTempUseCase) Execute(input string) (TempOutput, error) {
	cep, err := entity.NewCEP(string(input))
	if err != nil {
		return TempOutput{}, err
	}

	location, err := g.LocationClient.GetLocation(cep)
	if err != nil {
		return TempOutput{}, err
	}

	tempC, err := g.WeatherClient.GetWeather(location)
	if err != nil {
		return TempOutput{}, err
	}

	temp := entity.NewTemperature(tempC)

	dto := TempOutput{
		TempC: math.Round(temp.TempC*10) / 10,
		TempF: math.Round(temp.TempF*10) / 10,
		TempK: math.Round(temp.TempK*10) / 10,
	}

	return dto, nil
}
