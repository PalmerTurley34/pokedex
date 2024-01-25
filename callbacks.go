package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
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
	os.Exit(0)
	return nil
}

func mapCommand(cfg *config, parsedCmd []string) error {
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}
	fmt.Println("Next Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - Area: %v\n", area.Name)
	}
	cfg.nextLocationURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func mapbCommand(cfg *config, parsedCmd []string) error {
	if cfg.prevLocationAreaURL == nil {
		return fmt.Errorf("already at first page, cannot go back")
	}
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return nil
	}
	fmt.Println("Previous Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - Area: %v\n", area.Name)
	}
	cfg.nextLocationURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func exploreCommand(cfg *config, parsedCmd []string) error {
	if len(parsedCmd) != 2 {
		return fmt.Errorf("explore takes exactly one area name")
	}
	areaName := parsedCmd[1]
	resp, err := cfg.pokeapiClient.GetArea(areaName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %v\n", pokemon.Pokemon.Name)
	}
	return nil
}

func catchCommand(cfg *config, parsedCmd []string) error {
	if len(parsedCmd) != 2 {
		return fmt.Errorf("catch takes exactly one pokemon name")
	}
	pokemonName := parsedCmd[1]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at the %v...\n", pokemon.Name)
	const threshold = 60
	randNum := rand.Intn(pokemon.BaseExperience)
	time.Sleep(time.Second)
	if randNum > threshold {
		return fmt.Errorf("failed to catch the %v", pokemon.Name)
	}
	fmt.Printf("You caught a %v!\n", pokemon.Name)
	cfg.pokedex[pokemonName] = pokemon
	return nil
}

func inspectComand(cfg *config, parsedCmd []string) error {
	if len(parsedCmd) != 2 {
		return fmt.Errorf("inspect command takes exactly one pokemon name")
	}
	pokemonName := parsedCmd[1]
	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("you have not yet caught a %v. Catch one to add it to the pokedex", pokemonName)
	}
	fmt.Println("")
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf(" - %v\n", pokeType.Type.Name)
	}
	fmt.Println("")
	return nil
}
