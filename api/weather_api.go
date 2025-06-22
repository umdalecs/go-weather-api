package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/umdalecs/weather-api/config"
)

func RetrieveData(location string) (string, error) {
	baseUrl := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s", location)

	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", fmt.Errorf("error parsing url")
	}

	q := u.Query()
	q.Set("unitGroup", "metric")
	q.Set("key", config.Envs.ApiKey)
	q.Set("contentType", "json")
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return "", fmt.Errorf("error retrieving data from external api")
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading the external api response body")
	}

	return string(bodyBytes), nil
}
