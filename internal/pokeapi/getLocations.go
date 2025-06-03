package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetStructResponse(baseurl string) (PokeAPIJsonResponse, error) {

	url := baseurl + "/location-area"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeAPIJsonResponse{}, fmt.Errorf("couldn't get the requested location area")
	}

	resp, err := c.httpClient.Do(req)
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
