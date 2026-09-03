package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) FetchPokemon(locationName string) (ExplorePokemon, error) {
	url := baseURL + "/location-area/" + locationName

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExplorePokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ExplorePokemon{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ExplorePokemon{}, err
	}

	var pokemonList ExplorePokemon
	if err := json.Unmarshal(data, &pokemonList); err != nil {
		return ExplorePokemon{}, err
	}

	return pokemonList, nil
}
