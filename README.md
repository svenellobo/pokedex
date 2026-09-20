# Pokedex CLI

A command-line Pokédex written in Go. Interactive REPL that queries the
[PokéAPI](https://pokeapi.co) to explore locations, catch Pokémon and
inspect what you've caught.

Built as a guided project from the Boot.dev Go course.

## Install

    git clone https://github.com/svenellobo/[pokedex]
    cd [pokedex]
    go build -o pokedex
    go run .


Commands:

- `help` - list available commands
- `map` / `mapb` - page forward and back through location areas
- `explore <area>` - list Pokémon found in an area
- `catch <pokemon>` - attempt to catch a Pokémon
- `inspect <pokemon>` - show stats for a Pokémon you've caught
- `pokedex` - list everything you've caught
- `exit` - quit

## Notes

API responses are cached in memory with a configurable expiry. The cache is
safe for concurrent access via a mutex and reaped by a background goroutine.
