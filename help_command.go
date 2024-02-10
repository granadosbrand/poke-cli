package main

import "fmt"

func commandHelp(c *config, args ...string) error {

	fmt.Println("Poke-CLI usage: ")
	for _, command := range commands() {
		fmt.Printf("- %s : \n", command.name)
		fmt.Printf("%s \n", command.description)

	}
	return nil
}
