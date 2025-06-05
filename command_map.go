package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {

	locations, err := cfg.client.GetStructResponse(cfg.nextLocations)
	if err != nil {
		return err
	}

	cfg.nextLocations = &locations.Next
	cfg.previousLocations = &locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousLocations == nil {
		return errors.New("you're on the first page")
	}
	locations, err := cfg.client.GetStructResponse(cfg.previousLocations)
	if err != nil {
		return err
	}
	cfg.nextLocations = &locations.Next
	cfg.previousLocations = &locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil

}
