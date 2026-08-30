package main

import (
    "fmt"
    "errors"  
)


func commandMap(cfg *config) error {
    locationResp, err := cfg.pokeapiClient.FetchLocations(cfg.nextLocationURL)
    if err != nil {
        return err
    }

    cfg.nextLocationURL = locationResp.Next
    cfg.previousLocationURL = locationResp.Previous

    for _, loc := range locationResp.Results {
        fmt.Println(loc.Name)
    }
	
    return nil

}

func commandMapb(cfg *config) error {
    if cfg.previousLocationURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.FetchLocations(cfg.previousLocationURL)
    if err != nil {
        return err
    }

    cfg.nextLocationURL = locationResp.Next
    cfg.previousLocationURL = locationResp.Previous

    for _, loc := range locationResp.Results {
        fmt.Println(loc.Name)
    }
	
    return nil

}