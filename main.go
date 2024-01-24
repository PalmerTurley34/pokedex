package main

import (
	"fmt"

	"github.com/PalmerTurley34/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationURL *string
	prevLocationAreaURL * string
}

func main() {
	fmt.Print("") // code will not run unless these print statemenst are here
	fmt.Print("") 
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	fmt.Print("")
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		nextLocationURL: nil,
		prevLocationAreaURL: nil,
	}
	startREPL(&cfg)
}