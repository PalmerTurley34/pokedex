package main

import (
	"github.com/PalmerTurley34/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	prevLocationAreaURL *string
	pokedex             map[string]pokeapi.PokemonResponse
}

func main() {
	cfg := config{
		pokeapiClient:       pokeapi.NewClient(),
		nextLocationURL:     nil,
		prevLocationAreaURL: nil,
		pokedex:             make(map[string]pokeapi.PokemonResponse),
	}
	startREPL(&cfg)
}
