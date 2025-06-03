package main

import "github.com/salamulyon/PokeDex/internal/pokeapi"

func main() {
	client := pokeapi.NewClient()
	cfg := &config{
		client: client,
	}
	startRepl(cfg)
}
