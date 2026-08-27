package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
)

type config struct {
	commands map[string]cliCommand
	previous *string
	next     *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type nameResult struct {
	Name string `json:"name"`
}

type locationAreaJson struct {
	Next     *string      `json:"next"`
	Previous *string      `json:"previous"`
	Results  []nameResult `json:"results"`
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		command, ok := cfg.commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			if err := command.callback(cfg); err != nil {
				fmt.Println(err)
			}

		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	keys := make([]string, 0, len(cfg.commands))

	for k := range cfg.commands {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		v := cfg.commands[k]
		fmt.Printf("%s: %s\n", v.name, v.description)
	}

	return nil
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/?limit=20"
    if cfg.next != nil {
        url = *cfg.next
    }
    return fetchLocations(cfg, url)

}

func commandMapb(cfg *config) error {
	if cfg.previous == nil {
        fmt.Println("you're on the first page")
        return nil
    }
    return fetchLocations(cfg, *cfg.previous)
}


func fetchLocations(cfg *config, url string) error {
	res, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("fetching locations: %w", err)
    }
    defer res.Body.Close()

    data, err := io.ReadAll(res.Body)
    if err != nil {
        return err
    }

    var locations locationAreaJson
    if err := json.Unmarshal(data, &locations); err != nil {
        return err
    }

    cfg.previous = locations.Previous
    cfg.next = locations.Next

    for _, loc := range locations.Results {
        fmt.Println(loc.Name)
    }
    return nil
}




