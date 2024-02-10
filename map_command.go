package main

import (
	"fmt"
)

func mapCommand(cfg *config, args ...string) error {
	pokeapiClient := cfg.pokeapiClient


	res, err := pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaURL = res.Next
	cfg.previousLocationAreaURL = res.Previous

	fmt.Println("\nLocation Areas: ")

	for _, area := range res.Results {
		fmt.Printf(" - %s \n", area.Name)
	}
	return nil
}
