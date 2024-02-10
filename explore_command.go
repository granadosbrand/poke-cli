package main

import (
	"errors"
	"fmt"
)

func explore(c *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("No location area provided")
	}

	res, err := c.pokeapiClient.ListPokemonInArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Pokemons in %s:\n", args[0])
	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf("  - %s \n", pokemon.Pokemon.Name)
	}
	return nil
}
