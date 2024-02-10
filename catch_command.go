package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

func catch(c *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("No Pokémon name provided")
	}

	pokemon, err := c.pokeapiClient.CatchPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("  ¡A Wild %s Has Appeared! \n", args[0])
	fmt.Printf("  %s's basic experience: %v \n", args[0], pokemon.BaseExperience)
	fmt.Printf("  Catching %s..\n", args[0])

	// Generate a random number
	n, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		return err
	}

	// If the random number is less than the base experience, the catch fails
	if n.Int64() < int64(pokemon.BaseExperience) {
		fmt.Printf("  %s escaped!\n", args[0])
	} else {
		fmt.Printf("  %s was caught!\n", args[0])
		c.caugthPokemons[pokemon.Name] = pokemon
	}

	// for _, pokemon:= range c.caugthPokemons {
	// 	fmt.Println(pokemon.Name)
	// }

	return nil
}
