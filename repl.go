package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/svenellobo/pokedex/internal/pokeapi"
)

type config struct {
	commands            map[string]cliCommand
	pokeapiClient       pokeapi.Client
	pokedex             map[string]pokeapi.CaughtPokemon
	previousLocationURL *string
	nextLocationURL     *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type nameResult struct {
	Name string `json:"name"`
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
		args := []string{}

		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		command, exists := cfg.commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		} else {
			if err := command.callback(cfg, args...); err != nil {
				fmt.Println(err)
			}

		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 location areas",
			callback:    commandMapb,
		},

		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},

		"catch": {
			name: "catch <pokemon_name>",
			description: "Catch pokemon and add them to your pokedex",
			callback: commandCatch,
		},
	}
}
