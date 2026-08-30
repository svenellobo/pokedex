package pokeapi


import (
	"io"
	"net/http"
	"encoding/json"
)

func (c *Client)FetchLocations(pageURL *string) (LocationAreaJson, error) {    
	url := baseURL + "/location-area"

    if pageURL != nil {
        url = *pageURL
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

    var locationResp LocationAreaJson
    if err := json.Unmarshal(data, &locationResp); err != nil {
        return LocationAreaJson{}, err
    }

    return locationResp, nil

}















