package main

import (
	"fmt"
	"time"

	"github.com/granadosbrand/poke-cli/internal/pokeapi"
	"github.com/granadosbrand/poke-cli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	pokeCache               pokecache.Cache
	caugthPokemons          map[string]pokeapi.Pokemon
}

func commands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "	Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "  Exit poke-cli",
			callback:    exitCommand,
		},
		"map": {
			name:        "map",
			description: "  Displays the names of 20 location areas in the Pokémon world.",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "  Displays the previous 20 location areas in the Pokémon world.",
			callback:    mapbCommand,
		},
		"explore": {
			name:        "explore <area>",
			description: "  Display a list of pokemon in a given area",
			callback:    explore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "  Catch a Pokémon",
			callback:    catch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "  Lets you see the details of the Pokémons you have already caught.",
			callback:    inspectCallback,
		},
		"pokedex": {
			name:        "pokedex",
			description: "  List of the Pokémons you have already caught.",
			callback:    pokedexCallback,
		},
	}

}

func main() {

	cfg := config{
		pokeapiClient:  pokeapi.NewClient(5 * time.Minute),
		caugthPokemons: map[string]pokeapi.Pokemon{},
	}

	fmt.Println("¡Welcome to Poke-Cli!")
	fmt.Println("Type help to see available commands")
	newConsole(&cfg)

}
