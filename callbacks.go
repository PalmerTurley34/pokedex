package main

import (
	"fmt"
	"os"
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
	os.Exit(0)
	return nil
}

func mapCommand(cfg *config) error {
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationURL)
	if err != nil {
		// log.Fatal(err)
		return nil
	}
	fmt.Println("Next Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("Location: %v\n", area.Name)
	}
	cfg.nextLocationURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous 
	return nil
}

func mapbCommand(cfg *config) error {
if cfg.prevLocationAreaURL == nil {
		return fmt.Errorf("already at first page, cannot go back")
	}
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return nil
	}
	fmt.Println("Previous Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("Location: %v\n", area.Name)
	}
	cfg.nextLocationURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}