package main

import (
	"github.com/PalmerTurley34/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationURL *string
	prevLocationAreaURL * string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		nextLocationURL: nil,
		prevLocationAreaURL: nil,
	}
	startREPL(&cfg)
}