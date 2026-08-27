package main

func main() {
	cfg := config{
		commands: map[string]cliCommand{
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
		},
		previous: nil,
		next:     nil,
	}
	startRepl(&cfg)
}
