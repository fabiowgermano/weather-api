package entity

type LocationClientInterface interface {
	GetLocation(CEP) (string, error)
}

type WeatherClientInterface interface {
	GetWeather(string) (float64, error)
}
