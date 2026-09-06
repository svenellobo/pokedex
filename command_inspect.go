package main

import (
	"fmt"
	"errors"
)
func commandInspect(cfg *config, args ...string) error{
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	pok, exists := cfg.pokedex[name]

	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pok.Name)
	fmt.Printf("Height: %d\n", pok.Height)
	fmt.Printf("Weight: %d\n", pok.Weight)

	fmt.Println("Stats:")
	for _, s := range pok.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pok.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil


}