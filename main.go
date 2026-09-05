package main

import (
	"time"

	"github.com/svenellobo/pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, 300*time.Second)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: client,
		pokedex: make(map[string]pokeapi.CaughtPokemon),
	}
	startRepl(cfg)
}
