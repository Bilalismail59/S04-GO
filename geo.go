package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Coordinates struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

func getCoordFfromCity(city string) (Coordinates, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", city, token)
	resp, err := http.Get(url)
	if err != nil {
		return Coordinates{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Coordinates{}, err
	}

	var resultat []Coordinates
	if err := json.Unmarshal(body, &resultat); err != nil {
		return Coordinates{}, err
	}

	// Check if resultat is empty
	if len(resultat) == 0 {
		return Coordinates{}, fmt.Errorf("no coordinates found for city: %s", city)
	}

	return resultat[0], nil
}
