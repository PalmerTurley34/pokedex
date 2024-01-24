package main

import (
	"fmt"
	"os"

	"github.com/PalmerTurley34/pokedex/internal/pokeapi"
	// "log"
)

func helpCommand(cfg *config) error {
	fmt.Print("\nAvailable Commands:\n\n")
	for _, cmd := range getAllCommands() {
		line := fmt.Sprintf("%s: %s", cmd.name, cmd.desc)
		fmt.Println(line)
	}
	return nil
}

func exitCommand(cfg *config) error {
	fmt.Println("\nGoodbye!")
	cfg.pokeCache.Ticker.Stop()
	os.Exit(0)
	return nil
}

func mapCommand(cfg *config) (err error) {
	if cfg.nextLocationURL == nil {
		return fmt.Errorf("reached last page, cannot go forward")
	}
	respBytes, ok := cfg.pokeCache.Get(*cfg.nextLocationURL)
	if !ok {
		fmt.Println("fetching from api")
		respBytes, err = cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationURL)
		if err != nil {
			return nil
		}
	}
	resp, err := pokeapi.UnmarshalLocationAreas(respBytes)
	if err != nil {
		return err
	}
	fmt.Println("Next Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("Location: %v\n", area.Name)
	}
	cfg.pokeCache.Add(*cfg.nextLocationURL, respBytes)
	cfg.nextLocationURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous 
	return nil
}

func mapbCommand(cfg *config) (err error) {
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
		fmt.Printf("Location: %v\n", area.Name)
	}
	cfg.pokeCache.Add(*cfg.prevLocationAreaURL, respBytes)
	cfg.nextLocationURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}