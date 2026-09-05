package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const catchThreshold = 40

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		fmt.Printf("No %s named pokemon found\n", name)
		return nil
	}
	
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	baseXp := pokemonResp.BaseExperience
	var chance int
	if baseXp <= 0 {
		chance = 0
	} else {
		chance = rand.Intn(baseXp)
	}
	if chance > catchThreshold {
		fmt.Printf("%s escaped!\n", name)
	} else {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokedex[name] = pokemonResp
	}

	return nil
}
