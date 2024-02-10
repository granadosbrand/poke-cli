package main

import (
	"errors"
	"fmt"
)

func mapbCommand(cfg *config, args ...string) error {


	if cfg.previousLocationAreaURL == nil {
		return errors.New("We are in the first page")
	}


	pokeapiClient := cfg.pokeapiClient
	res, err := pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
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
