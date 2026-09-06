package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) < 1 {
		fmt.Println("You haven't caught any pokemon yet")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for k := range cfg.pokedex {
		fmt.Printf(" - %s\n", k)
	}

	return nil


} 