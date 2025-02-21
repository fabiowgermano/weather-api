package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/fabiowgermano/weather-api/internal/entity"
)

var ErrNotFoundZipcode = errors.New("can not find zipcode")

type LocationClient struct {
	BaseURL string
}

type LocationResponse struct {
	Localidade string `json:"localidade"`
	Erro       string `json:"erro"`
}

func NewLocationClient(url string) *LocationClient {
	return &LocationClient{BaseURL: url}
}

func (l *LocationClient) GetLocation(cep entity.CEP) (string, error) {
	url := strings.ReplaceAll(l.BaseURL, "@CEP", string(cep))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var location LocationResponse
	err = json.Unmarshal(body, &location)
	if err != nil {
		return "", err
	}

	if location.Erro == "true" {
		return "", ErrNotFoundZipcode
	}

	return location.Localidade, nil
}
