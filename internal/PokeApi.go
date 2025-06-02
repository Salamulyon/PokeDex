package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLocationArea(url string) (PokeAPIJsonResponse, error) {
	if url == nil{
		url := "https://pokeapi.co/api/v2/location-area/"
	}
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeAPIJsonResponse{}, fmt.Errorf("couldn't get the requested location area")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return PokeAPIJsonResponse{}, fmt.Errorf("client couldnt process the request")
	}

	defer resp.Body.Close()

	var pokeResponse PokeAPIJsonResponse
	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&pokeResponse); err != nil {
		return PokeAPIJsonResponse{}, fmt.Errorf("couldnt decode received JSON")
	}

	return pokeResponse, nil
}
