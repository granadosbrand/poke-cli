package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("Aiudaaa")
	return nil
}

func commandExit() error {
	fmt.Println("amonosss")
	return nil
}
func commands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "Help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "Exit",
			description: "Exit poke-cli",
			callback:    commandExit,
		},
	}

}

func console(commands map[string]cliCommand) {

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("poke-cli > ")
		scanner.Scan() // Lee la próxima línea de input
		cmdString := scanner.Text()

		switch cmdString {
		case "help":
			fmt.Println("Usage: ")
			for _, command := range commands {
				fmt.Println(command.name, command.description)
			}
		case "exit":
			return

		default:
			fmt.Println("Undefined command")

		}

	}

}

func main() {
	fmt.Println("¡Welcome to Poke-Cli!")
	console(commands())
}
