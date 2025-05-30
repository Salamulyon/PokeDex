package main

import (
	"fmt"
)

func commandHelp() error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println()
	for _, value := range commands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}

	return nil
}
