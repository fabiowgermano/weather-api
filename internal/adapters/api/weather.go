package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type WeatherClient struct {
	BaseURL string
}

type WeatherRequest struct {
	Locations []LocationRequest `json:"locations"`
}

type LocationRequest struct {
	Q string `json:"q"`
}

type WeatherResponse struct {
	Bulk []BulkResponse `json:"bulk"`
}

type BulkResponse struct {
	Query QueryResponse `json:"query"`
}

type QueryResponse struct {
	Current CurrentResponse `json:"current"`
}

type CurrentResponse struct {
	TempC float64 `json:"temp_c"`
}

func NewWeatherClient(url string, apiKey string) *WeatherClient {
	baseUrl := strings.ReplaceAll(url, "@APIKEY", apiKey)
	return &WeatherClient{BaseURL: baseUrl}
}

func (w *WeatherClient) GetWeather(city string) (float64, error) {
	data := WeatherRequest{
		Locations: []LocationRequest{
			{Q: city},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPost, w.BaseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return 0, err
	}

	return weather.Bulk[0].Query.Current.TempC, nil
}
