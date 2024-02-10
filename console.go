package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(str string) []string {
	wordsSlice := strings.Fields(strings.ToLower(str))
	return wordsSlice
}

func newConsole(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("poke-cli > ")
		scanner.Scan()
		cmdString := scanner.Text()

		cleaned := cleanInput(cmdString)

		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]
		args := []string{}

		if len(command) > 1 {
			args = cleaned[1:]
		}

		availableCommands, ok := commands()[command]

		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := availableCommands.callback(cfg, args...)

		if err != nil {
			fmt.Println(err)
		}

	}

}
