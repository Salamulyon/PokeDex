package main

import "github.com/Salamulyon/Pokedex/internal/pokeapi"

//import "github.com/Salamulyon/PokeDex/internal/pokeapi"

func main() {
	pokeClient := pokeapi.NewClient()
	cfg := &config{
		client: pokeClient,
	}
	startRepl(cfg)
}
