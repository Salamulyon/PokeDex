package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = make(map[string]cliCommand)

func cleanInput(text string) []string {

	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl() {

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exits the Pokedex",
		callback:    commandExit,
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		// prints without new line
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) < 1 {
			continue
		}
		command := words[0]
		if value, ok := commands[command]; ok {
			err := value.callback()
			fmt.Println(err)
		} else {
			fmt.Println("Unknown command")
		}
	}

}

func commandExit() error {

	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp() error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, value := range commands {
		fmt.Printf("%s: %s\n", value, value.description)
	}
	return nil
}
