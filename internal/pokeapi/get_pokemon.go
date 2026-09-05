package pokeapi

import (
	"encoding/json"	
	"io"
	"net/http"
)

func(c *Client) GetPokemon(pokemonName string) (CaughtPokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if pok, exists := c.cache.CacheGet(url); exists {
		var pokemon CaughtPokemon
		if err := json.Unmarshal(pok, &pokemon); err != nil{
			return CaughtPokemon{}, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CaughtPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CaughtPokemon{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return CaughtPokemon{}, err
	}

	var pokemon CaughtPokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {		
		return CaughtPokemon{}, err	
	}

	c.cache.CacheAdd(url, data)
	return pokemon, nil
}