package main

import (
	"errors"
	"fmt"
)

func inspectCallback(c *config, args ...string) error {

	pokemonName := args[0]

	pokemon, ok := c.caugthPokemons[pokemonName]
	if !ok {
		return errors.New("You have not caught that pokémon")
	}

	fmt.Printf("Pokemon details: \n")
	fmt.Printf("	Name: %s", pokemon.Name)
	fmt.Printf("	Height: %v", pokemon.Height)
	fmt.Printf("	Weight: %v", pokemon.Weight)
	fmt.Printf("	Experience: %v\n", pokemon.BaseExperience)
	fmt.Printf("	Stats: \n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("		%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("	Types: \n")
	for _, t := range pokemon.Types {
		fmt.Printf("		%s\n", t.Type.Name)
	}

	return nil

}
