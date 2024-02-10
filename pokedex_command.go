package main

import "fmt"

func pokedexCallback(c *config, args ...string) error {

	fmt.Println("Caught Pokemons: ")
	if len(c.caugthPokemons) == 0 {
		fmt.Println("  You have not caught any pokémon yet")
	}
	for _, pokemon := range c.caugthPokemons {
		fmt.Println("  " ,pokemon.Name)
	}

	return nil
}
