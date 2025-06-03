package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Salamulyon/PokeDex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	client            *pokeapi.Client
	nextLocations     *string
	previousLocations *string
}

var commands = make(map[string]cliCommand)

func cleanInput(text string) []string {

	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandsInit() {

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Get the next page of locations",
		callback:    commandMapf,
	}
	commands["mapb"] = cliCommand{
		name:        "map",
		description: "Get the previous page of locations",
		callback:    commandMapb,
	}
}

func startRepl(cfg *config) {

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
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}

}
