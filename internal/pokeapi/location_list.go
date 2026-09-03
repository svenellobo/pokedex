package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) FetchLocations(pageURL *string) (LocationAreaJson, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if ch, ok := c.cache.CacheGet(url); ok {
		var locationResp LocationAreaJson
		if err := json.Unmarshal(ch, &locationResp); err != nil {
			return LocationAreaJson{}, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaJson{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaJson{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaJson{}, err
	}

	c.cache.CacheAdd(url, data)

	var locationResp LocationAreaJson
	if err := json.Unmarshal(data, &locationResp); err != nil {
		return LocationAreaJson{}, err
	}

	return locationResp, nil

}
