package main

import (
	"fmt"
	"os"

	"github.com/PalmerTurley34/pokedex/internal/pokeapi"
	// "log"
)

func helpCommand(cfg *config, parsedCmd []string) error {
	fmt.Print("\nAvailable Commands:\n\n")
	for _, cmd := range getAllCommands() {
		line := fmt.Sprintf("%s: %s", cmd.name, cmd.desc)
		fmt.Println(line)
	}
	return nil
}

func exitCommand(cfg *config, parsedCmd []string) error {
	fmt.Println("\nGoodbye!")
	cfg.pokeCache.Ticker.Stop()
	os.Exit(0)
	return nil
}

func mapCommand(cfg *config, parsedCmd []string) (err error) {
	if cfg.nextLocationURL == nil {
		return fmt.Errorf("reached last page, cannot go forward")
	}
	respBytes, ok := cfg.pokeCache.Get(*cfg.nextLocationURL)
	if !ok {
		respBytes, err = cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationURL)
		if err != nil {
			return err
		}
	}
	resp, err := pokeapi.UnmarshalLocationAreas(respBytes)
	if err != nil {
		return err
	}
	fmt.Println("Next Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - Area: %v\n", area.Name)
	}
	cfg.pokeCache.Add(*cfg.nextLocationURL, respBytes)
	cfg.nextLocationURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func mapbCommand(cfg *config, parsedCmd []string) (err error) {
	if cfg.prevLocationAreaURL == nil {
		return fmt.Errorf("already at first page, cannot go back")
	}
	respBytes, ok := cfg.pokeCache.Get(*cfg.prevLocationAreaURL)
	if !ok {
		respBytes, err = cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationAreaURL)
		if err != nil {
			return nil
		}
	}
	resp, err := pokeapi.UnmarshalLocationAreas(respBytes)
	if err != nil {
		return err
	}
	fmt.Println("Previous Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - Area: %v\n", area.Name)
	}
	cfg.pokeCache.Add(*cfg.prevLocationAreaURL, respBytes)
	cfg.nextLocationURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func exploreCommand(cfg *config, parsedCmd []string) (err error) {
	if len(parsedCmd) != 2 {
		return fmt.Errorf("explore takes exactly one area name")
	}
	areaName := parsedCmd[1]
	respBytes, ok := cfg.pokeCache.Get(areaName)
	if !ok {
		respBytes, err = cfg.pokeapiClient.GetArea(areaName)
		if err != nil {
			return err
		}
	}
	resp, err := pokeapi.UnmarshalArea(respBytes)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
	}
	cfg.pokeCache.Add(areaName, respBytes)
	return nil
}
