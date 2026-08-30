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
	previousLocationURL *string
	nextLocationURL     *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
	}
}
