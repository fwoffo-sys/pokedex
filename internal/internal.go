package internal

import (
	"io"
	"net/http"
	"encoding/json"
)

type LocationPage struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetPokeLocations(url string) (LocationPage, error) {
	locationPage := LocationPage{}
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationPage, err
	}
	res, err := client.Do(req)
	if err != nil {
		return locationPage, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationPage, err
	}
	if err = json.Unmarshal(data, &locationPage); err != nil {
		return locationPage, err
	}
	return locationPage, nil
}