package main

import (
	"fmt"
	"sort"
)

func commandHelp(cfg *config, args ...string) error {
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
