package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

type cliCommand struct {
	name          string
	description   string
	callback      func() error
	configuration *config
}

type config struct {
	Next     url.URL
	Previous url.URL
}

var commands = make(map[string]cliCommand)

func cleanInput(text string) []string {

	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandsInit() {

	commands["help"] = cliCommand{
		name:          "help",
		description:   "Displays a help message",
		callback:      commandHelp,
		configuration: &config{},
	}
	commands["exit"] = cliCommand{
		name:          "exit",
		description:   "Exit the Pokedex",
		callback:      commandExit,
		configuration: &config{},
	}
	commands["map"] = cliCommand{
		name:          "map",
		description:   "Display the names of 20 locations in the world",
		callback:      commandMap,
		configuration: &config{},
	}

}

func startRepl() {

	commandsInit()

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
		if command, ok := commands[command]; ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}

}
